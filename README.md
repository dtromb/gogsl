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
                
                cd dtromb/gogsl; go generate


    This step is necessary because the datatype sizes on various machines will differ, and the bindings need to be generated based on these. This necessarily means code will not be portable across architectures (the alternative would have been to use per-binding size-matched types, however Go would then require explicit casts when these were used!)
    
    If you are on an **x86-64** platform, you can skip this step as the generated code that ships with the project is for this architecture.
    
*   Import and use the bindings:
    
                import "github.com/dtromb/gogsl/<package>"

    The bindings very closely follow the API, you should be able to use them in your code simply using the  [GSL Manual](https://www.gnu.org/software/gsl/manual) as a reference.
    
---

## Memory management, other "gotchas"

Not-insignificant dissonances occur between the call patterns and data use cases in the C library and the Go bindings.  Resolving these was the primary difficulty that need to be overcome, to create the proper "smoothness".  In some cases end users need to be aware of these issues.   They will be summarized here:

* No need to clean up...

Go is a garbage-collected language, and we attempt to take full advantage of this facility.   GSL kindly uses the *_[c]alloc(), *_free() pattern to lifecycle most of its managed-memory objects, and that allocation pattern is mapped to the bindings - there are respectively *(A|Ca)lloc() and *automated object cleanup* available.  Any functions created with an alloc call do not need the corresponding free in Go - simply let these fall unreachable and the corresponding GSL-managed memory will be freed.

* ...but careful to not make a mess.

Now the bad:  these objects sometimes reference each other in non-obvious ways, and callers must in some cases stop referencing "live" objects if others are no longer reachable.  This isn't as horrible as it sounds - to be safe there are just a couple of rules to follow.   First, for `vector` and `matrix` view types, make sure that the underlying object (Go slice/array, **gogsl** object, etc) is not collected when the view is still active. Secondly, do not in any case store references to the special objects returned by functions with a trailing underscore, such as `GslVector`'s `Slice_()` operation - use these only as function call arguments to "cast" values (for example, maybe you have a vector, but the function you want to pass its data to requires a Go slice - use `Slice_()`, but do not keep the slice around for later).   

These limitations exist because of the reasonable (in C, anyway) assumption on the part of GSL that you will be casting pointers between types and passing them around, as well as using other C-like "memory block" features.  Right now breaking these "rules" will get you a segfault, or worse.   It's a TODO to make this a bit nicer - holding Go-side references to objects "in use" by the C side memory, and using the `GslErrorHandler` mechanism to recoverably fail instead of  signaling.  See the `/<package>/test` directories for proper use patterns.

---

## License

The code is made available under the version 2 GNU Public License (Copyleft).  If you use it, you must release your source with your product and follow certain other restrictions.   See the `LICENSE` file in the distribution for details.

To say hello, or if you have any questions (please report issues on github) <dave.trombley@gmail.com>.   I'd love to know who you are and what you're using gogsl for!

