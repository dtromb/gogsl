package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

/*
	#include <stdlib.h>

	// XXX - We assume here that size_t is unsigned - is this valid everywhere?
	int get_c_size_t_size() {
	int c;
	size_t k;
		k = k - 1;
		while (k) {
			k >>= 1;
			c += 1;
		}
		return c;
	}

	int get_c_short_size() {
	int c;
	unsigned short k;
		k = k - 1;
		while (k) {
			k >>= 1;
			c += 1;
		}
		return c;
	}

	int get_c_int_size() {
	int c;
	unsigned int k;
		k = k - 1;
		while (k) {
			k >>= 1;
			c += 1;
		}
		return c;
	}

	int get_c_long_size() {
	int c;
	unsigned long k;
		k = k - 1;
		while (k) {
			k >>= 1;
			c += 1;
		}
		return c;
	}

	int get_c_long_long_size() {
	int c;
	unsigned long long k;
		k = k - 1;
		while (k) {
			k >>= 1;
			c += 1;
		}
		return c;
	}

*/
import "C"

var FUNCTION_TOK_RE *regexp.Regexp = regexp.MustCompile(`([a-zA-Z0-9_]+\s*|\*|,|\[\]|\(|\))`)
var METADATA_RE *regexp.Regexp = regexp.MustCompile(`#\s*([a-zA-Z0-9_]+)\s*\{(\s*((("([^"\\]|\\")*")|\[\s*(("([^"\\]|\\")*"\s*),?)*\])\s*,?\s*)*\s*)\}`)
var METADATA_ARGS_RE *regexp.Regexp = regexp.MustCompile(`(("([^"\\]|\\")*")|\[\s*(("([^"\\]|\\")*"\s*),?)*\])\s*`)
var QUOTED_REP *regexp.Regexp = regexp.MustCompile(`\\"`)

type Metainfo struct {
	key  string
	args [][]string
}

type Wrapper struct {
	preCall       []string
	callReplace   string
	postCall      []string
	returnReplace string
	suppressCall  bool
}

type TypeMapStrategy int

const (
	VALUE_CAST TypeMapStrategy = iota
	GSL_COMPLEX
	GSL_COMPLEX_FLOAT
	GSL_REFERENCE
	GSL_REFERENCE_DTOR
	GSL_FUNCTION
	GSL_FUNCTION_FDF
	GSL_MONTE_FUNCTION
	OUTPTR
	SLICEPTR
	SLICEARRAY
	FILE
	STRING
	VOID_RETURN
	CPTR
)

func TypeMapStrategyFromString(name string) (TypeMapStrategy, error) {
	switch strings.ToLower(name) {
	case "value_cast":
		return VALUE_CAST, nil
	case "outptr":
		return OUTPTR, nil
	case "gsl_complex":
		return GSL_COMPLEX, nil
	case "gsl_complex_float":
		return GSL_COMPLEX_FLOAT, nil
	case "gsl_reference":
		return GSL_REFERENCE, nil
	case "gsl_reference_dtor":
		return GSL_REFERENCE_DTOR, nil
	case "gsl_function":
		return GSL_FUNCTION, nil
	case "gsl_function_fdf":
		return GSL_FUNCTION_FDF, nil
	case "sliceptr":
		return SLICEPTR, nil
	case "slicearray":
		return SLICEARRAY, nil
	case "file":
		return FILE, nil
	case "string":
		return STRING, nil
	case "void_return":
		return VOID_RETURN, nil
	case "cptr":
		return CPTR, nil
	}
	return TypeMapStrategy(-1), errors.New("Invalid type map strategy: '" + name + "'")
}

func RestyleCNameToGo(name string, upcase bool) string {
	parts := strings.Split(name, "_")
	for i, part := range parts {
		if part == "" {
			continue
		}
		parts[i] = strings.ToUpper(part[0:1]) + part[1:]
	}
	k := strings.Join(parts, "")
	if !upcase {
		k = strings.ToLower(k[0:1]) + k[1:]
	}
	return k
}

//# PACKAGE{"gogsl",["gsl_math.h"],["GSL_","gsl_"]}
func ParseMetainfo(line string) (*Metainfo, error) {
	match := METADATA_RE.FindStringSubmatch(line)
	if match == nil {
		return nil, errors.New("invalid metainfo line: " + line)
	}
	name := match[1]
	args := strings.Trim(match[2], " \t")
	if args == "" {
		return &Metainfo{key: name}, nil
	}
	//argParts := strings.Split(args,",")
	argParts := METADATA_ARGS_RE.FindAllString(args, -1)
	md := &Metainfo{key: name, args: make([][]string, len(argParts))}
	for i, arg := range argParts {
		arg = strings.Trim(arg, " \t")
		if arg[0] == '[' {
			arg = arg[1 : len(arg)-1]
			arg = strings.Trim(arg, " \t")
			if arg != "" {
				argParts := strings.Split(arg, ",")
				md.args[i] = make([]string, len(argParts))
				for j, qp := range argParts {
					qp = strings.Trim(qp, " \t")
					qp = qp[1 : len(qp)-1]
					qp = QUOTED_REP.ReplaceAllString(qp, "\"")
					md.args[i][j] = qp
				}
			} else {
				md.args[i] = []string{}
			}
		} else {
			arg = arg[1 : len(arg)-1]
			arg = QUOTED_REP.ReplaceAllString(arg, "\"")
			md.args[i] = make([]string, 1)
			md.args[i][0] = arg
		}
	}
	return md, nil
}

func StartSection(section string) {
	fmt.Println("// SECTION: " + section)
}

func GetIntSize() int {
	var c, k uint
	k = k - 1
	for k > 0 {
		c += 1
		k >>= 1
	}
	return int(c)
}

func GetCSizeTSize() int {
	return int(C.get_c_size_t_size())
}

func GetCShortSize() int {
	return int(C.get_c_short_size())
}

func GetCIntSize() int {
	return int(C.get_c_int_size())
}

func GetCLongSize() int {
	return int(C.get_c_long_size())
}

func GetCLongLongSize() int {
	return int(C.get_c_long_long_size())
}

var global_GO_IMPORTS []string = []string{}
var global_STRIP_PREFIXES []string = []string{}
var global_OUTPUT *ByteBufferWriter
var global_PREAMBLE *ByteBufferWriter
var global_PACKAGE_PATH string
var global_IS_COMPLEX_PACKAGE bool

type ByteBufferWriter struct {
	buf []byte
}

func (bbw *ByteBufferWriter) Write(p []byte) (n int, err error) {
	bbw.buf = append(bbw.buf, p...)
	return len(p), nil
}

func (bbw *ByteBufferWriter) String() string {
	return string(bbw.buf)
}

func AddImport(name string) {
	fmt.Println("IMPORT! " + CurrentPackageName() + " <- " + name)
	path := filepath.Clean(global_PACKAGE_PATH)
	basePath := filepath.Clean(global_PACKAGE_BASE_PATH)
	if strings.HasPrefix(path, basePath) {
		path = path[len(basePath):]
	}
	if strings.HasPrefix(path, "/") {
		path = path[1:]
	}
	if path == name {
		return
	}
	//fmt.Println("CURRENT PACKAGE: "+path)
	for _, imp := range global_GO_IMPORTS {
		if imp == name {
			return
		}
	}
	global_GO_IMPORTS = append(global_GO_IMPORTS, name)
	if strings.HasSuffix(name, "complex") {
		global_PREAMBLE.Write([]byte("import complex_ \"" + name + "\"\n"))
	} else {
		global_PREAMBLE.Write([]byte("import \"" + name + "\"\n"))
	}
}

var global_CTYPES_TO_REF_TYPES map[string]string = make(map[string]string)
var global_PACKAGES_BY_REFERENCE_TYPE map[string]string = make(map[string]string)
var global_CTYPES_TO_ENUM_TYPES map[string]string = make(map[string]string)
var global_PACKAGES_BY_ENUM_TYPE map[string]string = make(map[string]string)

func GslReferenceType(name string, fields []string) {
	AddImport("unsafe")
	AddImport("github.com/dtromb/gogsl")
	goTypeName := RestyleCNameToGo(name, true)
	typedecl := fmt.Sprintf("type %s struct {\n   gogsl.GslReference\n", goTypeName)
	for _, field := range fields {
		typedecl = fmt.Sprintf("%s   %s\n", typedecl, field)
	}
	typedecl = fmt.Sprintf("%s}\n\n", typedecl)
	global_OUTPUT.Write([]byte(typedecl))
	global_PACKAGES_BY_REFERENCE_TYPE[name] = global_PACKAGE_PATH
	global_CTYPES_TO_REF_TYPES[name] = goTypeName
	fmt.Println("REF TYPE: " + name + " -> " + goTypeName)
}

var global_FUNCTION_NAME_OVERRIDE string

func GslName(name string) {
	global_FUNCTION_NAME_OVERRIDE = name
}

func GslEnumType(name string) {
	var rname string
	if strings.HasSuffix(name, "_t") {
		rname = name[0 : len(name)-2]
	} else {
		rname = name
	}
	goTypeName := RestyleCNameToGo(rname, true)
	typedecl := fmt.Sprintf("type %s uint32\n", goTypeName)
	global_OUTPUT.Write([]byte(typedecl))
	global_PACKAGES_BY_ENUM_TYPE[name] = global_PACKAGE_PATH
	global_CTYPES_TO_ENUM_TYPES[name] = goTypeName
	fmt.Println("ENUM TYPE: " + name + " -> " + goTypeName)
}

func StartPackage(name string, includes []string, prefixes []string) {
	if global_PACKAGE_PATH != "" {
		err := FinishPackage()
		if err != nil {
			panic(err.Error())
		}
	}
	fmt.Println("// PACKAGE: " + name)
	global_IS_COMPLEX_PACKAGE = strings.HasSuffix(name, "/complex")
	// Reset type mappings to default.
	global_CTYPE_TO_GOTYPE = make(map[string]string)
	global_CTYPE_TO_STRATEGY = make(map[string]TypeMapStrategy)
	fmt.Println("RESET TYPES")
	global_CTYPE_TO_GOTYPE["double"] = "float64"
	global_CTYPE_TO_GOTYPE["float"] = "float32"
	global_CTYPE_TO_GOTYPE["size_t"] = global_SIZE_T_TYPE
	global_CTYPE_TO_GOTYPE["int"] = global_INT_TYPE
	global_CTYPE_TO_GOTYPE["short"] = global_SHORT_TYPE
	global_CTYPE_TO_GOTYPE["unsigned short"] = global_SHORT_TYPE
	global_CTYPE_TO_GOTYPE["unsigned int"] = global_UINT_TYPE
	global_CTYPE_TO_GOTYPE["long"] = global_LONG_TYPE
	global_CTYPE_TO_GOTYPE["unsigned long"] = global_ULONG_TYPE
	global_CTYPE_TO_GOTYPE["long long"] = global_LONG_LONG_TYPE
	global_CTYPE_TO_GOTYPE["unsigned long long"] = global_ULONG_LONG_TYPE
	global_CTYPE_TO_GOTYPE["char"] = "int8"
	global_CTYPE_TO_GOTYPE["unsigned char"] = "uint8"
	global_CTYPE_TO_GOTYPE["gsl_complex"] = "complex128"
	global_CTYPE_TO_GOTYPE["gsl_complex_packed_ptr"] = "complex128"
	global_CTYPE_TO_GOTYPE["gsl_complex_packed_array"] = "complex128"
	global_CTYPE_TO_GOTYPE["gsl_complex_float"] = "complex64"
	global_CTYPE_TO_GOTYPE["gsl_complex_float_packed_ptr"] = "complex64"
	global_CTYPE_TO_GOTYPE["gsl_complex_float_packed_array"] = "complex64"
	global_CTYPE_TO_GOTYPE["gsl_function"] = "*GslFunction"
	global_CTYPE_TO_GOTYPE["gsl_function*"] = "*GslFunction"
	global_CTYPE_TO_GOTYPE["gsl_function_fdf"] = "*GslFunctionFdf"
	global_CTYPE_TO_GOTYPE["gsl_function_fdf*"] = "*GslFunctionFdf"
	// XXX - Need to introspect machine to valid the above two types!
	global_CTYPE_TO_GOTYPE["FILE"] = "os.File"
	// Invert the type map
	global_GOTYPE_TO_CTYPE = make(map[string]string)
	for k, v := range global_CTYPE_TO_GOTYPE {
		global_GOTYPE_TO_CTYPE[v] = k
	}
	path := filepath.Join(global_PACKAGE_BASE_PATH, name)
	global_PACKAGE_PATH = path
	fmt.Println("// PATH: " + path)
	global_STRIP_PREFIXES = make([]string, len(prefixes))
	copy(global_STRIP_PREFIXES, prefixes)
	for _, include := range includes {
		fmt.Println("// INCLUDE:  #include <" + include + ">")
	}
	packageParts := strings.Split(name, "/")
	packageName := packageParts[len(packageParts)-1]
	global_OUTPUT = &ByteBufferWriter{}
	global_PREAMBLE = &ByteBufferWriter{}
	global_PREAMBLE.Write([]byte("//////////        AUTOMATICALLY GENERATED CODE - DO NOT EDIT        //////////\n"))
	global_PREAMBLE.Write([]byte(fmt.Sprintf("package %s\n\n/*\n   #cgo pkg-config: --define-variable=prefix=. gsl\n", packageName)))
	for _, include := range includes {
		global_PREAMBLE.Write([]byte(fmt.Sprintf("   #include <%s>\n", include)))
	}
	global_PREAMBLE.Write([]byte("*/\nimport \"C\"\n\n"))
}

func QualifyPackage(pkg string, symbol string) string {
	fmt.Println("QUALIFY: " + pkg + " | " + CurrentPackageName() + ": " + symbol)
	if CurrentPackageName() == pkg {
		return symbol
	}
	pkgPath := strings.Split(pkg, "/")
	return fmt.Sprintf("%s.%s", pkgPath[len(pkgPath)-1], symbol)
}

func QualifyType(cType string, goType string) string {
	var baseType string
	if strings.HasSuffix(cType, "*") {
		baseType = cType[0 : len(cType)-1]
	}
	if enumPkg, has := global_PACKAGES_BY_ENUM_TYPE[cType]; has {
		pkg := CanonicalPackageName(enumPkg)
		AddImport(pkg)
		goType = QualifyPackage(pkg, goType)
		return goType
	}
	if enumPkg, has := global_PACKAGES_BY_ENUM_TYPE[baseType]; has {
		pkg := CanonicalPackageName(enumPkg)
		AddImport(pkg)
		goType = QualifyPackage(pkg, goType)
		return goType
	}
	if structPkg, has := global_PACKAGES_BY_REFERENCE_TYPE[cType]; has {
		pkg := CanonicalPackageName(structPkg)
		AddImport(pkg)
		goType = QualifyPackage(pkg, goType)
		return goType
	}
	if structPkg, has := global_PACKAGES_BY_REFERENCE_TYPE[baseType]; has {
		pkg := CanonicalPackageName(structPkg)
		AddImport(pkg)
		goType = QualifyPackage(pkg, goType)
		return goType
	}
	return goType
}

func CanonicalPackageName(path string) string {
	path = filepath.Clean(path)
	base := filepath.Clean(global_PACKAGE_BASE_PATH)
	if path[0] == '/' {
		path = path[1:]
	}
	if base[0] == '/' {
		base = base[1:]
	}
	if strings.HasSuffix(path, "/") {
		path = path[0 : len(path)-1]
	}
	if strings.HasSuffix(base, "/") {
		base = base[0 : len(base)-1]
	}
	if strings.HasPrefix(path, base) {
		path = path[len(base):]
	}
	if path[0] == '/' {
		path = path[1:]
	}
	fmt.Printf("PACKAGE CAN: %s | %s\n", path, base)
	return path
}

func CurrentPackageName() string {
	return CanonicalPackageName(global_PACKAGE_PATH)
}

func FinishPackage() error {
	err := os.MkdirAll(global_PACKAGE_PATH, 0700)
	if err != nil {
		return err
	}
	pathParts := strings.Split(global_PACKAGE_PATH, "/")
	fName := pathParts[len(pathParts)-1] + ".go"
	fPath := filepath.Join(global_PACKAGE_PATH, fName)
	fmt.Println("WRITING " + fPath + "...")
	fileOut, err := os.OpenFile(fPath, os.O_TRUNC|os.O_RDWR|os.O_CREATE, 0700)
	if err != nil {
		return err
	}
	fileOut.WriteString(global_PREAMBLE.String())
	fileOut.Write([]byte("\n"))
	fileOut.WriteString(global_OUTPUT.String())
	err = fileOut.Close()
	if err != nil {
		return err
	}
	global_PACKAGE_PATH = ""
	global_OUTPUT = &ByteBufferWriter{}
	global_PREAMBLE = &ByteBufferWriter{}
	global_GO_IMPORTS = []string{}
	return nil
}

func GetTypeMapStrategy(cType string, place int) (TypeMapStrategy, error) {
	if strat, has := global_ARG_MAP_STRATEGIES[place]; has {
		return strat, nil
	}
	if strat, has := global_CTYPE_TO_STRATEGY[cType]; has {
		return strat, nil
	}
	switch cType {
	case "gsl_complex":
		return GSL_COMPLEX, nil
	case "gsl_complex_float":
		return GSL_COMPLEX_FLOAT, nil
	case "gsl_function":
		return GSL_FUNCTION, nil
	case "gsl_function*":
		return GSL_FUNCTION, nil
	case "gsl_function_fdf":
		return GSL_FUNCTION_FDF, nil
	case "gsl_function_fdf*":
		return GSL_FUNCTION_FDF, nil
	case "void":
		return VOID_RETURN, nil
	}
	var refName string
	if strings.HasSuffix(cType, "*") {
		refName = cType[0 : len(cType)-1]
	} else {
		refName = cType
	}
	if strings.HasSuffix(cType, "[]") {
		return SLICEARRAY, nil
	}
	if _, has := global_CTYPES_TO_REF_TYPES[refName]; has {
		return GSL_REFERENCE, nil
	}
	if _, has := global_CTYPES_TO_REF_TYPES[cType]; has {
		return GSL_REFERENCE, nil
	}
	switch cType {
	case "gsl_complex_packed_ptr":
		return SLICEPTR, nil
	case "gsl_complex_float_packed_ptr":
		return SLICEPTR, nil
	case "gsl_complex_packed_array":
		return SLICEPTR, nil
	case "gsl_complex_float_packed_array":
		return SLICEPTR, nil
	case "FILE":
		fallthrough
	case "FILE*":
		return FILE, nil
	}
	return VALUE_CAST, nil
}

var global_CTYPE_TO_GOTYPE map[string]string = make(map[string]string)
var global_GOTYPE_TO_CTYPE map[string]string = make(map[string]string)
var global_CTYPE_TO_CGOTYPE map[string]string = make(map[string]string)
var global_CTYPE_TO_STRATEGY map[string]TypeMapStrategy = make(map[string]TypeMapStrategy)

func AddTypemap(ctype string, gotype string, cgotype string) {
	global_CTYPE_TO_GOTYPE[ctype] = gotype
	global_GOTYPE_TO_CTYPE[gotype] = ctype
	if cgotype != "" {
		global_CTYPE_TO_CGOTYPE[ctype] = cgotype
	}
	fmt.Printf("TYPEMAP %s %s %s\n", ctype, gotype, cgotype)
	fmt.Printf("TYPEMAP len=%d\n", len(global_CTYPE_TO_GOTYPE))
}

func AddTypemapWithStrategy(ctype string, gotype string, cgotype string, tmStrat string) error {
	AddTypemap(ctype, gotype, cgotype)
	tms, err := TypeMapStrategyFromString(tmStrat)
	if err != nil {
		return err
	}
	global_CTYPE_TO_STRATEGY[ctype] = tms
	return nil
}

// Maps C type name to correct Go type.
func MapCTypeName(cName string) (string, error) {
	if goType, has := global_CTYPE_TO_GOTYPE[cName]; has {
		return goType, nil
	}
	fmt.Println("NOTYPE: " + cName)
	fmt.Printf("TYPEMAP len=%d\n", len(global_CTYPE_TO_GOTYPE))
	var refName string
	if strings.HasSuffix(cName, "*") {
		refName = cName[0 : len(cName)-1]
	} else {
		refName = cName
	}
	if refType, has := global_CTYPES_TO_REF_TYPES[refName]; has {
		return refType, nil
	}
	if enumType, has := global_CTYPES_TO_ENUM_TYPES[refName]; has {
		return enumType, nil
	}
	if strings.HasSuffix(cName, "[]") {
		baseName := cName[0 : len(cName)-2]
		goBaseType, err := MapCTypeName(baseName)
		if err == nil {
			return "[]" + goBaseType, nil
		}
	}
	return "?", errors.New("cannot map C type to Go type: " + cName)
}

func IsPointerType(typename string) bool {
	switch typename {
	case "gsl_complex_packed_ptr":
		return true
	case "gsl_complex_float_packed_ptr":
		return true
	case "gsl_complex_packed_array":
		return true
	case "gsl_complex_float_packed_array":
		return true
	}
	return strings.HasSuffix(typename, "*")
}

func ValueCastExpression(castType string, expr string) string {
	var castExpr string
	if strings.HasPrefix(castType, "*") ||
		strings.HasPrefix(castType, "[]") ||
		strings.HasPrefix(castType, "map") ||
		strings.HasPrefix(castType, "chan") {
		castExpr = fmt.Sprintf("(%s)", castType)
	} else {
		castExpr = castType
	}
	return fmt.Sprintf("%s(%s)", castExpr, expr)
}

func ConstructFunctionWrapper(cName string, cRetType []string, cArgTypes []string, cArgNames []string) (string, error) {
	var buf []byte
	var body Wrapper
	var goRetType string
	var err error
	var goArgTypes []string
	var argMapStrategies []TypeMapStrategy = make([]TypeMapStrategy, len(cArgTypes))
	var outptrReturnIndex map[int]int = make(map[int]int)
	var returnReferenceType string

	var returnMapStrategies []TypeMapStrategy = []TypeMapStrategy{}
	for _, rt := range cRetType {
		rms, err := GetTypeMapStrategy(rt, 0)
		if err == nil {
			returnMapStrategies = append(returnMapStrategies, rms)
		} else {
			returnMapStrategies = append(returnMapStrategies, VALUE_CAST)
		}
	}

	// First write function signature.
	cName = strings.Trim(cName, " \t")
	callName := "C." + cName
	for _, prefix := range global_STRIP_PREFIXES {
		if strings.HasPrefix(cName, prefix+"_") {
			cName = cName[len(prefix):]
			break
		}
	}
	cName = strings.Trim(cName, " \t")
	if global_FUNCTION_NAME_OVERRIDE == "" {
		cName = RestyleCNameToGo(cName, true)
	} else {
		cName = global_FUNCTION_NAME_OVERRIDE
		global_FUNCTION_NAME_OVERRIDE = ""
	}
	if len(cArgTypes) != len(cArgNames) {
		return "", errors.New("arg name/type length mismatch")
	}

	// Special gsl reference dtor handling per default.
	if strings.HasSuffix(callName, "_free") && len(cArgTypes) == 1 && len(cRetType) == 1 {
		var argType string
		if strings.HasSuffix(cArgTypes[0], "*") {
			argType = cArgTypes[0][0 : len(cArgTypes[0])-1]
		} else {
			argType = cArgTypes[0]
		}
		if rType, has := global_CTYPES_TO_REF_TYPES[argType]; has && cRetType[0] == "void" {
			returnMapStrategies[0] = GSL_REFERENCE_DTOR
			returnReferenceType = rType
		}
	}

	// Write the function decl.
	if returnMapStrategies[0] == GSL_REFERENCE_DTOR {
		buf = append(buf, fmt.Sprintf("func (x *%s) Dispose(", returnReferenceType)...)
	} else {
		buf = append(buf, fmt.Sprintf("func %s(", cName)...)
	}

	for i, arg := range cArgNames {
		arg = RestyleCNameToGo(arg, false)
		if i == 0 && returnMapStrategies[0] == GSL_REFERENCE_DTOR {
			continue
		}
		tms, err := GetTypeMapStrategy(cArgTypes[i], i+1)
		argMapStrategies[i] = tms
		if err != nil {
			return "", err
		}
		switch tms {
		case CPTR:
			fallthrough
		case VALUE_CAST:
			{
				var goArgType string
				var err error
				if overrideType, has := global_ARG_MAP_TYPE_OVERRIDES[i]; has {
					goArgType = overrideType
				} else {
					goArgType, err = MapCTypeName(cArgTypes[i])
				}
				goArgTypes = append(goArgTypes, goArgType)
				if err != nil {
					return "", err
				}
				if enumPkg, has := global_PACKAGES_BY_ENUM_TYPE[cArgTypes[i]]; has {
					pkg := CanonicalPackageName(enumPkg)
					AddImport(pkg)
					goArgType = QualifyPackage(pkg, goArgType)
				}
				buf = append(buf, fmt.Sprintf("%s %s", arg, goArgType)...)
				if i < len(cArgNames)-1 {
					buf = append(buf, ", "...)
				}
			}
		case GSL_COMPLEX:
			{
				goArgTypes = append(goArgTypes, "complex128")
				buf = append(buf, fmt.Sprintf("%s complex128", arg)...)
				if i < len(cArgNames)-1 {
					buf = append(buf, ", "...)
				}
			}
		case GSL_COMPLEX_FLOAT:
			{
				goArgTypes = append(goArgTypes, "complex64")
				buf = append(buf, fmt.Sprintf("%s complex64", arg)...)
				if i < len(cArgNames)-1 {
					buf = append(buf, ", "...)
				}
			}
		case GSL_FUNCTION:
			{
				AddImport("github.com/dtromb/gogsl")
				typeName := QualifyPackage("github.com/dtromb/gogsl", "GslFunction")
				goArgTypes = append(goArgTypes, typeName)
				buf = append(buf, fmt.Sprintf("%s *%s", arg, typeName)...)
				if i < len(cArgNames)-1 {
					buf = append(buf, ", "...)
				}
			}
		case GSL_FUNCTION_FDF:
			{
				AddImport("github.com/dtromb/gogsl")
				typeName := QualifyPackage("github.com/dtromb/gogsl", "GslFunctionFdf")
				goArgTypes = append(goArgTypes, typeName)
				buf = append(buf, fmt.Sprintf("%s *%s", arg, typeName)...)
				if i < len(cArgNames)-1 {
					buf = append(buf, ", "...)
				}
			}
		case GSL_MONTE_FUNCTION:
			{
				AddImport("github.com/dtromb/gogsl")
				typeName := QualifyPackage("github.com/dtromb/gogsl", "GslMonteFunction")
				goArgTypes = append(goArgTypes, typeName)
				buf = append(buf, fmt.Sprintf("%s *%s", arg, typeName)...)
				if i < len(cArgNames)-1 {
					buf = append(buf, ", "...)
				}
			}
		case GSL_REFERENCE:
			{
				var refType string
				if strings.HasSuffix(cArgTypes[i], "*") {
					refType = cArgTypes[i][0 : len(cArgTypes[i])-1]
				} else {
					refType = cArgTypes[i]
				}
				pkg := CanonicalPackageName(global_PACKAGES_BY_REFERENCE_TYPE[refType])
				fmt.Println("PBRT[" + refType + "] = " + pkg)
				AddImport(pkg)
				if rType, has := global_CTYPES_TO_REF_TYPES[refType]; has {
					buf = append(buf, fmt.Sprintf("%s *%s", arg, QualifyPackage(pkg, rType))...)
					if i < len(cArgNames)-1 {
						buf = append(buf, ", "...)
					}
				} else {
					return "", errors.New("undefined reference type: '" + refType + "'")
				}
			}
		case OUTPTR:
			{
				AddImport("unsafe")
				if !strings.HasSuffix(cArgTypes[i], "*") {
					return "", errors.New("OUTPTR target type must be a pointer")
				}
				retType := cArgTypes[i][0 : len(cArgTypes[i])-1]
				outptrReturnIndex[len(cRetType)] = i
				cRetType = append(cRetType, retType)
				cGoType, err := MapCTypeNameToCGo(retType)
				if err != nil {
					return "", err
				}
				returnMapStrategies = append(returnMapStrategies, OUTPTR)
				body.preCall = append(body.preCall, fmt.Sprintf("var _outptr_%d %s", i, cGoType))
			}
		case SLICEPTR:
			{
				var refType string
				if !IsPointerType(cArgTypes[i]) {
					return "", errors.New("SLICEPTR argument type is not a pointer")
				}
				if strings.HasSuffix(cArgTypes[i], "*") {
					refType = cArgTypes[i][0 : len(cArgTypes[i])-1]
				} else {
					refType = cArgTypes[i]
				}
				baseGoType, err := MapCTypeName(refType)
				if err != nil {
					return "", err
				}
				buf = append(buf, fmt.Sprintf("%s []%s", arg, baseGoType)...)
				if i < len(cArgNames)-1 {
					buf = append(buf, ", "...)
				}
			}
		case SLICEARRAY:
			{

				var refType string
				if strings.HasSuffix(cArgTypes[i], "[]") {
					refType = cArgTypes[i][0 : len(cArgTypes[i])-2]
				} else {
					return "", errors.New("SLICEARRAY argument type is not an array")
				}
				baseGoType, err := MapCTypeName(refType)
				if err != nil {
					return "", err
				}
				buf = append(buf, fmt.Sprintf("%s []%s", arg, baseGoType)...)
				if i < len(cArgNames)-1 {
					buf = append(buf, ", "...)
				}
			}
		case FILE:
			{
				AddImport("os")
				buf = append(buf, fmt.Sprintf("%s *os.File", arg)...)
				if i < len(cArgNames)-1 {
					buf = append(buf, ", "...)
				}
			}
		case STRING:
			{
				buf = append(buf, fmt.Sprintf("%s string", arg)...)
				if i < len(cArgNames)-1 {
					buf = append(buf, ", "...)
				}
			}
		default:
			{
				return "", errors.New("unknown argument type map strategy")
			}
		}
	}
	voidResult := cRetType[0] == "void"
	dangleChk := string(buf)
	if strings.HasSuffix(dangleChk, ", ") {
		buf = []byte(dangleChk[0 : len(dangleChk)-2])
	}

	switch len(cRetType) {
	case 0:
		{
			buf = append(buf, ") { \n"...)
		}
	case 1:
		{
			fmt.Println("RET TYPE: " + cRetType[0])
			if cRetType[0] != "void" {
				var err error
				if overrideType, has := global_ARG_MAP_TYPE_OVERRIDES[0]; has {
					goRetType = overrideType
				} else {
					goRetType, err = MapCTypeName(cRetType[0])
				}
				if err != nil {
					return "", err
				}
				fmt.Printf("QUALIFY RETURN TYPE %s | %s\n", cRetType[0], goRetType)
				goRetType = QualifyType(cRetType[0], goRetType)
				fmt.Printf("     ... %s\n", goRetType)
				if returnMapStrategies[0] == GSL_REFERENCE {
					buf = append(buf, fmt.Sprintf(") *%s {\n", goRetType)...)
				} else {
					buf = append(buf, fmt.Sprintf(") %s {\n", goRetType)...)
				}
			} else {
				buf = append(buf, ") {\n"...)
			}
		}
	default:
		{
			buf = append(buf, ") ("...)
			for i, rt := range cRetType {
				if rt == "void" {

					continue
				}
				var gort string
				var err error
				if overrideType, has := global_ARG_MAP_TYPE_OVERRIDES[0]; i == 0 && has {
					gort = overrideType
				} else {
					gort, err = MapCTypeName(rt)
				}

				//gort, err := MapCTypeName(rt)
				if i == 0 {
					goRetType = gort
				}
				gort = QualifyType(rt, gort)
				if err != nil {
					return "", err
				}
				if returnMapStrategies[i] == GSL_REFERENCE {
					buf = append(buf, "*"+gort...)
				} else {
					buf = append(buf, gort...)
				}
				if i < len(cRetType)-1 {
					buf = append(buf, ", "...)
				}
			}
			buf = append(buf, ") {\n"...)
		}
	}

	// Transform the argument names.
	argNames := make([]string, len(cArgNames))
	for i, arg := range cArgNames {
		argNames[i] = RestyleCNameToGo(arg, false)
	}

	// Map the arguments to subexpressions.
	argExprs := make([]string, len(cArgNames))
	for i, arg := range argNames {
		if i == 0 && returnMapStrategies[0] == GSL_REFERENCE_DTOR {
			continue
		}
		tms, err := GetTypeMapStrategy(cArgTypes[i], i+1)
		if err != nil {
			return "", err
		}
		switch tms {
		case VALUE_CAST:
			{
				fmt.Printf("VALUE CAST: %d named %s as %s\n", i, argNames[i], cArgTypes[i])
				cgoType, err := MapCTypeNameToCGo(cArgTypes[i])
				if err != nil {
					return "", err
				}
				argExprs[i] = ValueCastExpression(cgoType, argNames[i])
			}
		case CPTR:
			{
				AddImport("unsafe")
				cgoType, err := MapCTypeNameToCGo(cArgTypes[i])
				if err != nil {
					return "", err
				}
				expr := fmt.Sprintf("unsafe.Pointer(%s.CPtr())", argNames[i])
				argExprs[i] = ValueCastExpression(cgoType, expr)
			}
		case GSL_COMPLEX:
			{
				AddImport("github.com/dtromb/gogsl/complex")
				if global_IS_COMPLEX_PACKAGE {
					body.preCall = append(body.preCall, fmt.Sprintf("_arg_%d := GoComplexToGsl(%s)", i, arg))
				} else {
					body.preCall = append(body.preCall, fmt.Sprintf("_arg_%d := complex_.GoComplexToGsl(%s)", i, arg))
				}
				argExprs[i] = fmt.Sprintf("*(*C.gsl_complex)(unsafe.Pointer(_arg_%d))", i)
			}
		case GSL_COMPLEX_FLOAT:
			{
				AddImport("github.com/dtromb/gogsl/complex")
				if global_IS_COMPLEX_PACKAGE {
					body.preCall = append(body.preCall, fmt.Sprintf("_arg_%d := GoComplexFloatToGsl(%s)", i, arg))
				} else {
					body.preCall = append(body.preCall, fmt.Sprintf("_arg_%d := complex_.GoComplexFloatToGsl(%s)", i, arg))
				}
				argExprs[i] = fmt.Sprintf("*(*C.gsl_complex_float)(unsafe.Pointer(_arg_%d))", i)
			}
		case GSL_FUNCTION:
			{
				AddImport("unsafe")
				fInit := QualifyPackage("github.com/dtromb/gogsl", "InitializeGslFunction")
				body.preCall = append(body.preCall, fmt.Sprintf("%s(%s)", fInit, arg))
				argExprs[i] = fmt.Sprintf("(*C.gsl_function)(unsafe.Pointer(%s.CPtr()))", arg)
			}
		case GSL_FUNCTION_FDF:
			{
				AddImport("unsafe")
				fInit := QualifyPackage("github.com/dtromb/gogsl", "InitializeGslFunctionFdf")
				body.preCall = append(body.preCall, fmt.Sprintf("%s(%s)", fInit, arg))
				argExprs[i] = fmt.Sprintf("(*C.gsl_function_fdf)(unsafe.Pointer(%s.CPtr()))", arg)
			}
		case GSL_REFERENCE:
			{
				var refType string
				AddImport("unsafe")
				if strings.HasSuffix(cArgTypes[i], "*") {
					refType = cArgTypes[i][0 : len(cArgTypes[i])-1]
				} else {
					refType = cArgTypes[i]
				}
				if _, has := global_CTYPES_TO_REF_TYPES[refType]; has {
					cgoType, err := MapCTypeNameToCGo(refType)
					if err != nil {
						return "", err
					}
					argExprs[i] = fmt.Sprintf("(*%s)(unsafe.Pointer(%s.Ptr()))", cgoType, arg)
				} else {
					return "", errors.New("undefined reference type: '" + refType + "'")
				}
			}
		case OUTPTR:
			{
				argExprs[i] = fmt.Sprintf("&_outptr_%d", i)
			}
		case SLICEPTR:
			{
				AddImport("reflect")
				AddImport("unsafe")
				cgoArgType, err := MapCTypeNameToCGo(cArgTypes[i])
				if err != nil {
					return "", err
				}
				body.preCall = append(body.preCall, fmt.Sprintf("_slice_header_%d := (*reflect.SliceHeader)(unsafe.Pointer(&%s))", i, arg))
				argExprs[i] = fmt.Sprintf("(%s)(unsafe.Pointer(_slice_header_%d.Data))", cgoArgType, i)
			}
		case SLICEARRAY:
			{
				AddImport("reflect")
				AddImport("unsafe")
				cgoArgType, err := MapCTypeNameToCGo(cArgTypes[i])
				if err != nil {
					return "", err
				}
				body.preCall = append(body.preCall, fmt.Sprintf("_slice_header_%d := (*reflect.SliceHeader)(unsafe.Pointer(&%s))", i, arg))
				argExprs[i] = fmt.Sprintf("(%s)(unsafe.Pointer(_slice_header_%d.Data))", cgoArgType, i)
			}
		// XXX - There is no way to manage file modes here, but in gsl uses of FILE* only ever write (fprintf pattern)
		//       If stdout or similar is passed, there will be an illegal seek, which we happily ignore.
		case FILE:
			{
				AddImport("unsafe")
				AddImport("github.com/dtromb/gogsl")
				body.preCall = append(body.preCall, fmt.Sprintf("_file_%d := C.fdopen(C.dup(C.int(%s.Fd())),(*C.char)(unsafe.Pointer(gogsl.APPEND_ONLY.Ptr())))", i, arg))
				body.postCall = append(body.postCall, fmt.Sprintf("C.fclose(_file_%d)", i))
				argExprs[i] = fmt.Sprintf("_file_%d", i)
			}
		case STRING:
			{
				AddImport("unsafe")
				argExprs[i] = fmt.Sprintf("_string_%d", i)
				body.preCall = append(body.preCall, fmt.Sprintf("%s := C.CString(%s)", argExprs[i], arg))
				body.postCall = append(body.postCall, fmt.Sprintf("C.free(unsafe.Pointer(%s))", argExprs[i]))
			}
		default:
			{
				return "", errors.New("unknown type mapping strategy")
			}
		}
	}

	// Create the C call expression
	if len(cRetType) == 0 {
		buf = append(buf, fmt.Sprintf("   %s()\n", callName)...)
	} else {
		var tms TypeMapStrategy
		if goRetType == "" {
			tms = returnMapStrategies[0]
		} else {
			tms, err = GetTypeMapStrategy(cRetType[0], 0)
			if err != nil {
				return "", err
			}
		}
		var callExprBuf []byte
		callExprBuf = append(callExprBuf, fmt.Sprintf("%s(", callName)...)
		for i, argExpr := range argExprs {
			callExprBuf = append(callExprBuf, argExpr...)
			if i < len(argExprs)-1 {
				callExprBuf = append(callExprBuf, ", "...)
			}
		}
		callExprBuf = append(callExprBuf, ")"...)
		callExpr := string(callExprBuf)
		switch tms {
		case VALUE_CAST:
			{
				callExpr = ValueCastExpression(goRetType, callExpr)
			}
		case GSL_COMPLEX:
			{
				AddImport("github.com/dtromb/gogsl/complex")
				if global_IS_COMPLEX_PACKAGE {
					body.returnReplace = "GslComplexToGo(uintptr(unsafe.Pointer(&_result)))"
				} else {
					body.returnReplace = "complex_.GslComplexToGo(uintptr(unsafe.Pointer(&_result)))"
				}
			}
		case GSL_COMPLEX_FLOAT:
			{
				AddImport("github.com/dtromb/gogsl/complex")
				if global_IS_COMPLEX_PACKAGE {
					body.returnReplace = "GslComplexFloatToGo(uintptr(unsafe.Pointer(&_result)))"
				} else {
					body.returnReplace = "complex_.GslComplexFloatToGo(uintptr(unsafe.Pointer(&_result)))"
				}
			}
		case GSL_FUNCTION:
			{
				panic("gsl_function return type unsupported")
			}
		case GSL_FUNCTION_FDF:
			{
				panic("gsl_function_fdf return type unsupported")
			}
		case GSL_REFERENCE:
			{
				if strings.HasSuffix(cRetType[0], "*") {
					AddImport("github.com/dtromb/gogsl")
					body.preCall = append(body.preCall, fmt.Sprintf("_ref := %s", callExpr))
					body.preCall = append(body.preCall, fmt.Sprintf("_result := &%s{}", goRetType))
					body.postCall = append(body.postCall, fmt.Sprintf("gogsl.InitializeGslReference(_result, uintptr(unsafe.Pointer(_ref)))"))
					body.returnReplace = "_result"
					body.suppressCall = true
				} else {
					fmt.Println(cRetType[0])
					AddImport("github.com/dtromb/gogsl")
					body.preCall = append(body.preCall, fmt.Sprintf("_ref := %s", callExpr))
					body.preCall = append(body.preCall, fmt.Sprintf("_result := &%s{}", goRetType))
					body.preCall = append(body.preCall, fmt.Sprintf("_result.CData = make([]byte, unsafe.Sizeof(_ref))"))
					body.postCall = append(body.postCall, fmt.Sprintf("gogsl.InitializeGslStatic(_result, uintptr(unsafe.Pointer(&_ref)))"))
					body.returnReplace = "_result"
					body.suppressCall = true
				}
			}
		case GSL_REFERENCE_DTOR:
			{
				cgoRetType, err := MapCTypeNameToCGo(cArgTypes[0])
				if err != nil {
					return "", err
				}
				body.preCall = append(body.preCall, fmt.Sprintf("%s((%s)(unsafe.Pointer(x.Ptr())))", callName, cgoRetType))
				body.suppressCall = true
			}
		case VOID_RETURN:
			{
				break
			}
		case STRING:
			{
				// We assume returned char* are not allocators and the library manages
				// the memory pointed to!  If this is not the case somewhere, a different
				// map strategy needs to be implemented for those returns...
				callExpr = fmt.Sprintf("C.GoString(%s)", callExpr)
			}
		case CPTR:
			{
				return "", errors.New("CPTR return type map strategy is unsupported")
			}
		default:
			{
				return "", errors.New("unknown return type map strategy ")
			}
		}
		for _, pc := range body.preCall {
			buf = append(buf, fmt.Sprintf("   %s\n", pc)...)
		}
		if !voidResult && len(body.postCall) == 0 && len(cRetType) == 1 && body.returnReplace == "" && !body.suppressCall {
			buf = append(buf, fmt.Sprintf("   return %s\n", callExpr)...)
		} else {
			// Do the call on a separate line from the return.
			if !body.suppressCall {
				if voidResult {
					buf = append(buf, fmt.Sprintf("   %s\n", callExpr)...)
				} else {
					buf = append(buf, fmt.Sprintf("   _result := %s\n", callExpr)...)
				}
			}
			for _, pc := range body.postCall {
				buf = append(buf, fmt.Sprintf("   %s\n", pc)...)
			}
			if body.returnReplace != "" {
				if cRetType[0] != "void" {
					buf = append(buf, fmt.Sprintf("   return %s\n", body.returnReplace)...)
				}
			} else {
				if cRetType[0] != "void" || len(cRetType) > 1 {
					buf = append(buf, "   return "...)
					for i, crt := range cRetType {
						if i == 0 {
							if voidResult {
								continue
							}
							buf = append(buf, "_result"...)
						} else {
							rms := returnMapStrategies[i]
							switch rms {
							case VALUE_CAST:
								{
									return "", errors.New("value cast in return mapping with no value")
								}
							case OUTPTR:
								{
									gort, err := MapCTypeName(crt)
									if err != nil {
										return "", err
									}
									opName := fmt.Sprintf("_outptr_%d", outptrReturnIndex[i])
									//expr := ValueCastExpression(gort, opName)
									buf = append(buf, fmt.Sprintf("*(*%s)(unsafe.Pointer(&%s))", gort, opName)...)
									// complex128(_outptr_3)
									// *(*complex128)(unsafe.Pointer(&_outptr_3))
								}
							default:
								{
									return "", errors.New("unknown return strategy")
								}
							}
						}
						if i < len(cRetType)-1 {
							buf = append(buf, ", "...)
						}
					}
					buf = append(buf, "\n"...)
				}
			}
		}
	}
	buf = append(buf, "}\n\n"...)
	return string(buf), nil
}

var global_ARG_MAP_STRATEGIES map[int]TypeMapStrategy = make(map[int]TypeMapStrategy)

func SetMapStrategy(place int, strategy string) error {
	tms, err := TypeMapStrategyFromString(strategy)
	if err != nil {
		return err
	}
	global_ARG_MAP_STRATEGIES[place] = tms
	return nil
}

var global_ARG_MAP_TYPE_OVERRIDES map[int]string = make(map[int]string)

func SetMapType(place int, typename string) error {
	global_ARG_MAP_TYPE_OVERRIDES[place] = typename
	return nil
}

func MapCTypeNameToCGo(name string) (string, error) {
	switch name {
	case "size_t":
		return "C.size_t", nil
	case "char":
		return "C.char", nil
	case "unsigned char":
		return "C.uchar", nil
	case "short":
		return "C.short", nil
	case "unsigned short":
		return "C.ushort", nil
	case "int":
		return "C.int", nil
	case "unsigned int":
		return "C.uint", nil
	case "long":
		return "C.long", nil
	case "unsigned long":
		return "C.ulong", nil
	case "long long":
		return "C.longlong", nil
	case "unsigned long long":
		return "C.ulonglong", nil
	case "float":
		return "C.float", nil
	case "double":
		return "C.double", nil
	case "gsl_function":
		return "C.gsl_function", nil
	case "gsl_complex":
		return "C.gsl_complex", nil
	case "gsl_complex_packed_ptr":
		return "*C.double", nil
	case "gsl_complex_float":
		return "C.gsl_complex_float", nil
	case "gsl_complex_float_packed_ptr":
		return "*C.float", nil
	case "gsl_complex_packed_array":
		return "*C.double", nil
	case "gsl_complex_float_packed_array":
		return "*C.float", nil
	case "FILE":
		return "C.FILE", nil
	}
	if cgo, has := global_CTYPE_TO_CGOTYPE[name]; has {
		return cgo, nil
	}
	if _, has := global_CTYPES_TO_REF_TYPES[name]; has {
		return "C." + name, nil
	}
	if _, has := global_CTYPES_TO_ENUM_TYPES[name]; has {
		return "C." + name, nil
	}
	if strings.HasSuffix(name, "*") {
		mapName, err := MapCTypeNameToCGo(name[0 : len(name)-1])
		if err != nil {
			return "?", err
		}
		return "*" + mapName, nil
	}
	if strings.HasSuffix(name, "[]") {
		mapName, err := MapCTypeNameToCGo(name[0 : len(name)-2])
		if err != nil {
			return "?", err
		}
		// XXX - Deal with multi-dimensional arrays here.
		return "*" + mapName, nil
	}
	return "?", errors.New("No mapping to CGo for C type '" + name + "'")
}

func GetNumberType(signed bool, n int) (string, error) {
	intSize := GetIntSize()
	if intSize == n {
		if signed {
			return "uint", nil
		}
		return "int", nil
	}
	switch n {
	case 8:
		{
			if signed {
				return "int8", nil
			} else {
				return "byte", nil
			}
		}
	case 16:
		{
			if signed {
				return "int16", nil
			} else {
				return "uint16", nil
			}
		}
	case 32:
		{
			if signed {
				return "int32", nil
			} else {
				return "uint32", nil
			}
		}
	case 64:
		{
			if signed {
				return "int64", nil
			} else {
				return "uint64", nil
			}
		}
	}
	return "?", errors.New("no default mapping for numeric type available")
}

var global_SIZE_T_TYPE string
var global_INT_TYPE string
var global_UINT_TYPE string
var global_SHORT_TYPE string
var global_USHORT_TYPE string
var global_LONG_TYPE string
var global_ULONG_TYPE string
var global_LONG_LONG_TYPE string
var global_ULONG_LONG_TYPE string
var global_C16_TYPE string
var global_C32_TYPE string
var global_C64_TYPE string
var global_PACKAGE_BASE_PATH string

var global_TEMPLATE_ADDITION_LIST []string = []string{}
var global_TEMPLATE_EXCEPTION_LIST []string = []string{}
var global_TEMPLATE_STD_PREFIXES []string = []string{
	"float", "float", "Float",
	"int", "int", "Int",
	"uint", "unsigned int", "Uint",
	"long", "long", "Long",
	"ulong", "unsigned long", "Ulong",
	"short", "short", "Short",
	"ushort", "unsigned short", "Ushort",
	"char", "char", "Char",
	"uchar", "unsigned char", "Uchar",
	"complex", "gsl_complex", "Complex",
	"complex_float", "gsl_complex_float", "ComplexFloat",
	// XXX - complex long double type is unsupported (no corresponding Go type!)
}

func ExpandPrefixTemplate(template string) []string {
	var lines []string
	var prefixes []string
	var skip bool
	for i := 0; i < len(global_TEMPLATE_STD_PREFIXES); i += 3 {
		skip = false
		for _, x := range global_TEMPLATE_EXCEPTION_LIST {
			if x == global_TEMPLATE_STD_PREFIXES[i] {
				skip = true
				break
			}
		}
		if !skip {
			prefixes = append(prefixes, global_TEMPLATE_STD_PREFIXES[i:i+3]...)
		}
	}
	if len(global_TEMPLATE_ADDITION_LIST) > 0 {
		panic("template addition list unimplemented")
	}
	for i := 0; i < len(prefixes); i += 3 {
		prefixName := global_TEMPLATE_STD_PREFIXES[i]
		typeName := global_TEMPLATE_STD_PREFIXES[i+1]
		goPrefixName := global_TEMPLATE_STD_PREFIXES[i+2]
		expanded := strings.Replace(template, "$", prefixName, -1)
		expanded = strings.Replace(expanded, "@", typeName, -1)
		expanded = strings.Replace(expanded, "%", goPrefixName, -1)
		newLines := strings.Split(expanded, "\n")
		var skipLine bool
		for k, newLine := range newLines {
			if skipLine {
				skipLine = false
				continue
			}
			newLine = strings.Trim(newLine, " \n")
			if newLine == "" {
				continue
			}
			fmt.Println("TEMPLATE LINE: " + newLine)
			if newLine[0] == '#' {
				mi, err := ParseMetainfo(newLine)
				if err == nil {
					switch mi.key {
					case "TEMPLATE_EXCEPT":
						{
							if len(mi.args) != 1 {
								panic("#TEMPLATE_EXCEPT expects exactly one argument")
							}
							for _, exceptType := range mi.args[0] {
								fmt.Printf("TMPEXCEPT %s/%s: %s\n", exceptType, prefixName, newLines[k+1])
								if exceptType == prefixName {
									skipLine = true
									break
								}
							}
							continue
						}
					}
				}
			}
			fmt.Println("TEMPLATE PRINT: " + expanded)
			lines = append(lines, newLine)
		}
	}
	global_TEMPLATE_ADDITION_LIST = []string{}
	global_TEMPLATE_EXCEPTION_LIST = []string{}
	return lines
}

func ReadPrefixTemplate(lines []string) (string, int) {
	var templateBuf []byte
	for i, line := range lines {
		line = strings.Trim(line, " \t")
		if line == "" {
			continue
		}
		if line[0] == '#' {
			mi, err := ParseMetainfo(line)
			if err != nil {
				panic("invalid metainfo in template: " + err.Error())
			}
			if mi.key == "GSL_PREFIX_TEMPLATE_END" {
				return string(templateBuf), (i + 1)
			}
		}
		templateBuf = append(templateBuf, (line + "\n")...)
	}
	panic("no #GSL_PREFIX_TEMPLATE_END at end of template")
}

func main() {

	// Introspect the machine and find the correct mappings for the C int, long, and long long
	// types.
	var err error
	cSizeTSize := GetCSizeTSize()
	cShortSize := GetCShortSize()
	cIntSize := GetCIntSize()
	cLongSize := GetCLongSize()
	cLongLongSize := GetCLongLongSize()

	if cShortSize == 16 {
		global_C16_TYPE = "short"
	} else if cIntSize == 16 {
		global_C16_TYPE = "int"
	} else {
		global_C16_TYPE = "int16_t"
	}
	if cIntSize == 32 {
		global_C32_TYPE = "int"
	} else if cShortSize == 32 {
		global_C32_TYPE = "short"
	} else if cLongSize == 32 {
		global_C32_TYPE = "long"
	} else {
		global_C64_TYPE = "int32_t"
	}
	if cIntSize == 64 {
		global_C64_TYPE = "int"
	} else if cLongSize == 64 {
		global_C64_TYPE = "long"
	} else if cLongLongSize == 64 {
		global_C64_TYPE = "long long"
	} else {
		global_C64_TYPE = "int64_t"
	}

	if global_SIZE_T_TYPE, err = GetNumberType(false, cSizeTSize); err != nil {
		panic("Cannot map C size_t: " + err.Error())
	}

	if global_INT_TYPE, err = GetNumberType(true, cIntSize); err != nil {
		panic("Cannot map C int: " + err.Error())
	}
	if global_UINT_TYPE, err = GetNumberType(false, cIntSize); err != nil {
		panic("Cannot map C unsigned int: " + err.Error())
	}
	if global_SHORT_TYPE, err = GetNumberType(true, cShortSize); err != nil {
		panic("Cannot map C short: " + err.Error())
	}
	if global_USHORT_TYPE, err = GetNumberType(false, cShortSize); err != nil {
		panic("Cannot map C unsigned short: " + err.Error())
	}
	if global_LONG_TYPE, err = GetNumberType(true, cLongSize); err != nil {
		panic("Cannot map C long: " + err.Error())
	}
	if global_ULONG_TYPE, err = GetNumberType(false, cLongSize); err != nil {
		panic("Cannot map C unsigned long: " + err.Error())
	}
	if global_LONG_LONG_TYPE, err = GetNumberType(true, cLongLongSize); err != nil {
		panic("Cannot map C long long: " + err.Error())
	}
	if global_ULONG_LONG_TYPE, err = GetNumberType(false, cLongLongSize); err != nil {
		panic("Cannot map C unsigned long long: " + err.Error())
	}

	var insideTemplate bool
	var templateLines []string
	var templatePosition int

	gopath := strings.Split(os.Getenv("GOPATH"), ":")[0]
	global_PACKAGE_BASE_PATH = filepath.Join(gopath, "src")
	flistBytes, err := ioutil.ReadFile(filepath.Join(
		gopath, "src",
		"github.com/dtromb/gogsl",
		"function-list",
	))
	flist := string(flistBytes)
	if err != nil {
		panic("Could not read function manifest")
	}
	lines := strings.Split(flist, "\n")
	var postImport bool
	//for lineNo, line := range lines {
	var lineNo int
	var line string
	for lineNo < len(lines) || insideTemplate {

		// Read current line, template aware.
		if insideTemplate {
			fmt.Printf("TEMPLATE: %d/%d\n", templatePosition, len(templateLines))
			if templatePosition >= len(templateLines) {
				insideTemplate = false
				fmt.Println("Finished template.")
				continue
			}
			line = templateLines[templatePosition]
			templatePosition += 1
		} else {
			line = lines[lineNo]
			lineNo += 1
		}

		line = strings.Trim(line, " \n\t")
		if strings.HasPrefix(line, "//") {
			continue
		}
		fmt.Printf("LINE: %d/%d %s\n", lineNo, len(lines), line)
		lineNo = lineNo
		if line == "" {
			continue
		}
		if line[0] == '#' {
			// Process metadata instruction
			md, err := ParseMetainfo(line)
			if err != nil {
				panic(err)
			}
			switch md.key {
			case "SECTION":
				{
					// # SECTION{"Vectors and Matricies"}
					if len(md.args) != 1 || len(md.args[0]) != 1 {
						panic("#SECTION expects exactly one argument")
					}
					StartSection(md.args[0][0])
				}
			case "PACKAGE":
				{
					// # PACKAGE{"gogsl.block", ["gsl_block.h"], ["gsl_block_"]}
					if len(md.args) != 3 {
						fmt.Printf("%d\n", len(md.args))
						panic("#PACKAGE expects exactly three arguments")
					}
					StartPackage(md.args[0][0], md.args[1], md.args[2])
				}
			case "IMPORT":
				{
					if len(md.args) != 1 {
						panic("#IMPORT expects exactly one argument")
					}
					for _, imp := range md.args[0] {
						AddImport(imp)
					}
					postImport = true
				}
			case "NEXTMAP":
				{
					// # NEXTMAP{"2","OUTPTR"}
					if len(md.args) != 2 || len(md.args[0]) != 1 || len(md.args[1]) != 1 {
						panic("#NEXTMAP expects exactly two string arguments")
					}
					place, err := strconv.Atoi(md.args[0][0])
					if err != nil {
						panic("#NEXTMAP first argument must be a number")
					}
					err = SetMapStrategy(place, md.args[1][0])
					if err != nil {
						panic(err.Error())
					}
					continue
				}
			case "NEXTTYPE":
				{
					if len(md.args) != 2 || len(md.args[0]) != 1 || len(md.args[1]) != 1 {
						panic("#NEXTTYPE expects exactly two string arguments")
					}
					place, err := strconv.Atoi(md.args[0][0])
					if err != nil {
						panic("#NEXTTYPE first argument must be a number")
					}
					err = SetMapType(place, md.args[1][0])
					if err != nil {
						panic(err.Error())
					}
					continue
				}
			case "GSL_REFTYPE":
				{
					if len(md.args) != 2 || len(md.args[0]) != 1 {
						panic("#GSL_REFTYPE expects two arguments with the first a single typename")
					}
					GslReferenceType(md.args[0][0], md.args[1])
				}
			case "GSL_ENUM":
				{
					if len(md.args) != 1 || len(md.args[0]) != 1 {
						panic("#GSL_ENUM expects exactly one string argument")
					}
					GslEnumType(md.args[0][0])
				}
			case "NAME":
				{
					if len(md.args) != 1 || len(md.args[0]) != 1 {
						panic("#NAME expects exactly one string argument")
					}
					GslName(md.args[0][0])
				}
			case "GSL_PREFIX_TEMPLATE_START":
				{
					fmt.Println("Processing template...")
					if len(md.args) != 0 {
						if len(md.args) != 2 {
							panic("#GSL_PREFIX_TEMPLATE_START takes zero or two arguments")
						}
						global_TEMPLATE_ADDITION_LIST = md.args[0]
						global_TEMPLATE_EXCEPTION_LIST = md.args[1]
					}
					if insideTemplate {
						panic("nested prefix templates unsupported")
					}
					template, n := ReadPrefixTemplate(lines[lineNo:])
					lineNo += n
					templateLines = ExpandPrefixTemplate(template)
					global_TEMPLATE_ADDITION_LIST = []string{}
					global_TEMPLATE_EXCEPTION_LIST = []string{}
					insideTemplate = true
					templatePosition = 0
					fmt.Println("Expanding template...")
					fmt.Printf("%d lines -> %d lines\n", n, len(templateLines))
				}
			case "GSL_PREFIX_TEMPLATE_END":
				{
					panic("unexpected #GSL_PREFIX_TEMPLATE_END")
				}
			case "TEMPLATE_EXCEPT":
				{
					panic("unexpected #TEMPLATE_EXCEPT")
				}
			case "TYPEMAP":
				{
					if len(md.args) < 3 || len(md.args) > 4 || len(md.args[0]) != 1 ||
						len(md.args[1]) != 1 || len(md.args[2]) != 1 {
						panic("#TYPEMAP expects three or four string arguments")
					}
					if len(md.args) == 3 {
						AddTypemap(md.args[0][0], md.args[1][0], md.args[2][0])
					} else {
						err := AddTypemapWithStrategy(md.args[0][0], md.args[1][0], md.args[2][0], md.args[3][0])
						if err != nil {
							panic(err.Error())
						}
					}
				}
			default:
				{
					panic("unknown metadata instruction '" + md.key + "'")
				}
			}
			continue
		}

		// Parse a function for translation.
		if postImport {
			global_OUTPUT.Write([]byte("\n"))
			postImport = false
		}
		toks := FUNCTION_TOK_RE.FindAllString(line, -1)
		var argStart int
		for i := 0; i < len(toks); i++ {
			if toks[i] == "(" {
				argStart = i + 1
				break
			}
		}
		if argStart == 1 || toks[len(toks)-1] != ")" {
			panic("invalid function line: " + line)
		}
		argParts := toks[argStart:len(toks)]
		namePart := toks[argStart-2]
		rtypeParts := toks[0 : argStart-2]
		returnType, err := TypeFromTypeParts(rtypeParts)
		if err != nil {
			panic("invalid return type in decl '" + line + "': " + err.Error())
		}
		//fmt.Printf("function %s \n   returns %s\n", namePart, returnType)
		var argSubs [][]string
		var spos int
		for i, ap := range argParts {
			ap = strings.Trim(ap, " \t")
			if ap == "," || ap == ")" {
				argSlice := argParts[spos:i]
				spos = i + 1
				argSubs = append(argSubs, argSlice)
			}
		}
		args := make([]string, len(argSubs))
		argNames := make([]string, len(argSubs))
		for i, sub := range argSubs {
			if len(sub) == 0 {
				if len(argSubs) != 1 {
					panic(fmt.Sprintf("empty arg decl with multiple args"))
				}
				args = []string{}
				argNames = []string{}
				break
			}
			if len(sub) < 2 {
				panic("invalid arg decl '" + strings.Join(sub, " ") + "' in function decl: " + line)
			}
			var argSuffixCount int
			var arrayType bool
			if sub[len(sub)-1] == "[]" {
				arrayType = true
				argSuffixCount = 2
			} else {
				argSuffixCount = 1
			}
			argNames[i] = sub[len(sub)-argSuffixCount]
			args[i], err = TypeFromTypeParts(sub[0 : len(sub)-argSuffixCount])
			if err != nil {
				panic("invalid arg type in decl '" + strings.Join(sub, " ") + "' in function decl: " + line + ":" + err.Error())
			}
			if arrayType {
				args[i] = args[i] + "[]"
			}
		}
		/*
			fmt.Printf("   args ")
			for i, arg := range args {
				fmt.Printf("%s %s", arg, argNames[i])
				if i < len(args)-1 {
					fmt.Printf(", ")
				}
			}
			fmt.Println()
		*/
		body, err := ConstructFunctionWrapper(namePart, []string{returnType}, args, argNames)
		if err != nil {
			panic("Could not construct function wrapper for '" + namePart + "': " + err.Error())
		}
		global_OUTPUT.Write([]byte(body))
		fmt.Println(body)

		// Reset mapping strategies.
		global_ARG_MAP_STRATEGIES = make(map[int]TypeMapStrategy)
		global_ARG_MAP_TYPE_OVERRIDES = make(map[int]string)
	}
	FinishPackage()
}

func TypeFromTypeParts(parts []string) (string, error) {
	var typestr []byte
	var haveName bool
	//fmt.Println("PARSE TYPE")
	for _, rtp := range parts {
		rtp = strings.Trim(rtp, " \t")
		switch rtp {
		case "const":
			{
				continue
			}
		case "unsigned":
			{
				typestr = append(typestr, "unsigned "...)
			}
		case "*":
			{
				typestr = append(typestr, "*"...)
				if !haveName {
					return "?", errors.New("invalid type (* before type name)")
				}
			}
		default:
			{
				if haveName {
					return "?", errors.New("invalid type (two type names) (last typstr=" + rtp + ")")
				}
				haveName = true
				//fmt.Println("NAME: "+rtp)
				typestr = append(typestr, rtp...)
			}
		}
	}
	return string(typestr), nil
}
