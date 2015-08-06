# gogsl
### A (nearly) complete binding of the GNU Scientific Library for Go

___

## Using gogsl
To use the **gogsl** bindings in your program, you'll need to do the following:

* Install [libgsl](https://www.gnu.org/software/gsl/):
    
    You will need the development headers in place at `/usr/include/gsl`.  (This location can be changed in `generator.go`.)  Your Linux distribution probably makes this step easy for you - for example, for a typical Fedora system:

                yum install gsl-devel
                
    will suffice.
* Get the distribution:
        
                go get github.com/dtromb/gogsl

* Generate the bindings:
                
                cd dtromb/gogsl/generator; go build .; ./generator


    This step is necessary because the datatype sizes on various machines will differ, and the bindings need to be generated based on these. This necessarily means code will not be portable across architectures (the alternative would have been to use per-binding size-matched types, however Go would then require explicit casts when these were used!)
    
    If you are on an **x86-64** platform, you can skip this step as the generated code that ships with the project is for this architecture.
    
*   Import and use the bindings:
    
                import "github.com/dtromb/gogsl/<package>"

    The bindings very closely follow the API, you should be able to use them in your code simply using the  [GSL Manual](https://www.gnu.org/software/gsl/manual) as a reference.
    
---
