# go-tutorial

## words to remember
 - qualified omitted explicit brace bracket encompasses extent

## go sytanx

### syntax

- ebnf

            |   alternation
            ()  grouping
            []  option (0 or 1 times)
            {}  repetition (0 to n times)

-  go programes are constructed by linking together packages.

         SourceFile   = PackageClause ";" { ImportDecl ";"} { TopLevelDecl ";" } .

         TopLevelDecl = Declaration | FunctionDecl | MethodDecl .

         Declaration  = Constecl | TypeDecl | VarDecl .

-  a package in turn is constructed from one or more source files that together declare constants,types,variables and 
    functions belonging to  the packages and which are accessible in all files of the same package.Those elemants may be
    exported and used in another package.
-  a package clause begins each source file and defines the package to which the file belongs.
-  PackageClause      = "package" PackageName. 
   PackageName         = identifier.
-  the package name must not be the blank identifier.
-  a set of files sharing the same PackageName form the implemention  of the package. A implemention may require that
    all source files for a package inhabit the same  directory.
-  ImportDecl         = "import" ( ImportSpec | "(" { ImportSpec } ")" ).
   ImportSpec         = [ "."| PackageName ] ImportPath
   ImportPath         = string_lit .
-  the PackageName in ImportSpec is used in qualified identifier to access exported identifiers of the package within the importing
     source file. It is declared in the file block. If the packagename is omitted, it defaults to the identifier specified in the package
     clause of the imported package. If an explicit period appears instead of name, all the package's exported identifiers declaried in that
     package block will be declared in the importing source file's file block and must be accessed without a qualifier.

### block

-  a block is possibly empty sequence of declarations and statements within matching brace brackets.
-  Block = "{" StatementList "}"
    StatementList = { Statement ";" } .
-  In addition to explicit blocks in the source code, there are implicit blocks:
    * The universe block encompasses all Go source text.
    * Each package  has a package block  containing  all Go source text for that package.
    * Each file has a file block containing all Go source text in that file.
    * Each "if", "for"  and "switch" statement is considered to be in  its own implicit block.
    * Each clause in a "switch" or "select" statement acts as an implicit block.

- a declaration binds a non-blank identifier to a constant,type,variable,function,label,or package.Every identifier in a program must be
    declared.No identifier may be declared twice in the same block,and no identifier may be declared in both the file and package block.
-  The blank identifier may be used like any other identifier in a declaration,but it does not  introduce a binding and thus is not declared.
    In the package block,the identifier init   may only be used for init function declarations, and  like the blank identifier it does not introuce
     a new binding.
-  the scope of a declared identifier is the extent of source text in which the identifier denotes the specified constant,type,variable,function
    label,or package.
     *   the predeclared identifier is the universe block.
     *   the scope of an identifier denoting a constant, type,variable,or function(but not method) declared at top level (outside any funciton) is the package block.
     *   the scope of the package name of an imported packages if the file block of the file containing the import declaration.
     *   the scope of an identifier denoting a method receiver,function parameter, or result variable is the funciion body.
     *   the scope of a constant or variable identifier declared inside a function begins at  the end of ConstSpec or VarSpec(ShortVarDecl for shor variable
     *    declarations) and ends at the end of the innermost containing block.
     *   the scope of a type identifier declared inside a function begins at the identifier in the TypeSpec and ends at the end of the innermost containing block.

-  an identifier declared in a block maybe reclared in an inner block.
-  the package clauses is not a declaration; the package name does not appear in any scope. Its purpose is to identify the files belonging to the
    same package and to specify the default package name for import declarations.
-  labels are declared  by labeled statements and are used in the "break" , "continue" and "goto" statements. It is  illegal to define a label that is never used.
       in contrast to other identifiers, labels are not block scoped  and do not conflict with identifiers that are not labels. The scope of a label
       is the function in which it is declared and exclude the body of any nested function.
-  the blank identifier is represented by the underscore character _.

### Program initialization and execution
- Within a package, package-level variables are initialized in declaration order but after any variables they depend on.
  a package-level variable is considered ready for initialized  if it is not yet initialized and either has no initialization expression or its initialization expression
  has no dependencies on uninitialized variables.Initializaiton proceeds by repeatedly  initializing the next package-level variable that is earliest in declaration order and ready for
  initialization,until there are no variables ready for initialization.

- If any variables are still uninitialized when this process ends,those variables are part of one or more initialization cycles,and the program is not valid.
- The declaration order of variables declared in multiple files is determined by the order in which the files are presented to the compiler.
- variables may also be initialized using functions named init declared in the package block.Mutiple such function may be defined.
- If multiple packages import a package, the imported package  will be initialized only once.
- Package initialization - variable initialization and the invocation of init functions - happends in a single goroutine, sequentially,one package at a time.
- An init function may launch other goroutines, which can run coucurrently with the initializaiton code.However,the 
    initializaiton always sequences the init functions:it will not invoke the next one until the previous has finished.
- to eusure reproducible initialization behavior,build systems are encoraged to present mutiple files belonging to the same package in lexical file name order to 
  to a compiler.

- a complete program is created by linking a single,unimported package called the main package with all the packages it imports,transitively. The main package must have package name
  main and declare a function main that takes no auguments and returns no value.

- Program execution begins by initializing the main package and then invoking the function main. When that function invocation returns,the program exits.It does not wait for other(non-main) 
  goroutines to complete.


- The predeclared type error is the conventional interface for representing an error condition,with the nil value representing no error.

- Execution errors such as attempting to index an array out of bounds trigger a run-time panic,equivalent to call of the built-in function panic with a value of th e
  implementation-definedinterface type runtime.Error.


### declarations
## constant decalrations
-  a constant declaration binds a list of identifiers to the values of a list of constant expressiongs.
    ConstDecl   =   "const" ( ConstSpec ) | "(" {  ConstSpec ";" } ")" ).
    ConstSpec   =   IdentifierList  [ [ Type ] "=" ExpressionList ].
    IdentifierList  =   identifier { "," identifier }.
    ExpressionList  = Expression { "," Expression }.

- If the type is present, all constants take the type specified.if the type if omitted,the constants
  take the individual types of the corresponding expressions.
  If the expression values are untyped constants,the declared constants remain untyped and the constant
  identifiers denote the constant values.

- within a parenthesized const declaration list the expression list may be omitted from any but the 
  first declaration. Such an empty list is equivalent to the textual substitution of the first preceding
  non-empty expression list and its type if any.Omitting the list of expressions is therefore equivalent to repeating
  the previous list.

- the predeclared identifier iota represents successive untyped integer constants.It is reset to 0 whenever the 
  reserved word const appears in the source and increments after each ConstSpec.

- within a ExpressionList, the value of each iota is the same because it is only incremented after each ConstSpec.

-
# Type declarations

- A type declaration binds an identifier, the type name, to a new type that has the same underlying type as an existing type.
  the operations defined for the existing tpe are also defined for the new type.
  The new type is different from the existing type.

    TypeDecl    =   "type" ( TypeSpec | "(" { TypeSpec ";" } ")" ).
    TypeSpec    =   identifier  Type .

- The declared type does not inherit any methods bound to the existing type, the the method set of an interface type
   or of emements of a composite type remains unchanged.

-# Variable declarations

- A variable declaration creates one or more variables, binds corresponding identifiers to them,and gives each a type and an initial value.

    VarDecl     =   "var" ( VarSpec | "(" { VarSpec ";" } ")" ).
    VarSpec     =   IdentifierList ( Type [ "=" ExpressionList ] | "=" ExpressionList )

-  If a type is not present,each variable is given the type of the corresponding initialization value in the assignment.
  if that value is an untyped constant,it is first converted to its default type. it it is an untyped boolean value, it is first converted to type bool.
    the predeclared value nil cannot be used to initialize a variable with no explicit type.

# Function declarations

- A function declaration binds an identifier, the function name, to a function.

    FunctionDecl    =   "func" FunctionName ( Function | Signature ).
    FunctionName    = identifier .
    Function    = Signature FunctionBody.
    FunctionBody    = Block

- If the function's signature declares result parameters, the function body's statement list
    must end in a terminating statement.

- A function declaration may omit the body. function implemented outside Go, such as an assembly routine.

# Method declarations

- A method is a function with a receiver. A method declaration binds an identifier, the
    method name, to a method, and associates the method with the reciever's base type.

    MethodDecl  =   "func" Receiver MethodName ( Function | Signature ).
    Receiver    =   Parameters .

- The parameter section must declare a single non-variadic parameter,the receiver.
- Its type must be of the form T or (xing T) .The T is called the receiver base type;
    it must not be a pointer or interface type and it must be declared in the 
    same package as the method.

- A non-blank receiver identifier must be unique in the method signature. If the receiver's
    value is not referenced inside the body of the method, its identifier may be omitted in the 
    declaration. The same applies in general to parameters of functions and methods.

- The type of a method is the type of a function with the receiver as first argument.
    However, a function declared this way is not a method.


### Types
#  A type determines the set of values and operations specific to values of that type.

    Type    =   TypeName    | TypeLit   | "(" Type ")".
    TypeName    =   identifier  | QualifiedIdent.
    TypeLit =   ArrayType   | StructType    | PointerType   | FunctionType | InterfaceType
                    | SliceType | MapType | ChannelType.

- unnamed types are specified using a type literal, which composes a new type from existing types.
- The method set of corresponding pointer type * T is the set of all methods declared with receiver * T or T


- A boolean type represents the set of Boolean truth values denoted by the predeclared constants true and false.
    The predeclared boolean type if bool.
- The predeclared numeric types are:
    uint8,uint16,uint32, uint64
    int8,int16,int32,int64
    float32, float64
    complex64,complex128
    byte    // alias for uint8
    rune   // alias for Uint32

    uint either 32 or 64 bits
    int     same size as uint
    uintptr an unsigned integer large enough to store the uninterpreted bits of a pointer value

- To avoid portablility issues all numeric type are distinct except byte and rune.Conversions are required.

- A string type represents the set of string value. A string value is a sequence of bytes. Strings are immutable.
  built-in function len can get the length of a string. The length is a compile-time constant if the string is a constant.
  It is illegal to take the address of such an element; if s[i] is the i'th byte of a string ,&s[i] is invalid.

- An array is a numbered sequence of elements of a single type.

    ArrayType   = "[ ArrayLength ]" ElementType
    ArrayLength =   Expression
    ElementType =   Type.
- The length is part of the array's type.it must evaluate to a non-negative constant representable by a value of type int.
  built-in function(len) can be used to get length of array.

- Array may be composed to form multidimensional types.

- The value of an uninitialized slice is nil

    SliceType   =   "[" "]" ElementType.

- The lenght of a slice may change. A slice, once initialized, is always associated with an underlying array that
    holds its elements.

- A new, initalized slice value for a given element type T is made using the built-in funcitoin make.A slice created with make
    always allocates a new,hidden array to which the returned slice value refers.

- like arrays, slices are always one-dimensional but may be composed to construct higher-dimentsional objects.with slices of slices
  the inner lengths may vary dynamically.Moreover,the inner slices must be initialized individually.

- A struct 

    StructType      =       "struct" "{" { FieldDecl ";" }"}".
    FieldDecl       =   ( IdentifierList Type | AnonymousField) [ Tag ].
    AnonymousField  =   [ "*" ] TypeName.
    Tag     =   string_lit.

- A embedded type must be specified as a type name T or as a pointer to a non-interface type name * T. and T itself may not be a pointer type.

- The unqualified type name acts as the filed name.

    * T2 //field name T2
    P.T3 //field name T3
    * P.T4 //field name T4

- Prometed fields act like ordinary fields of a struct.

- Given a struct type S and a type named T, promoted methods are include in the method set of the struct as follows:
    If S contains an anonymous field T, the method sets of S and * S both include promoted methods with receiver T.
    The method set of *S also includes promoted methods with receiver *T

    If S contains an anonymous field *T, the method sets of S and *S both include promoted methods with receiver T or *T.

- A pointer type denotes the set of all pointers to variables of a givern type, called the base type of the pointer.The value
  of an uninitialized pointer is nil

  PointerType   =   " * "  BaseType
  BaseType  =   Type.

- A function type denotes the set of all functions with the same parameter and result types.

    FunctionType    =   "func"  Signature.
    Signature   =   Parameters [ Result ]
    Result  =   Parameters | Type .
    Parameters  =   "( [ ParameterList [ "," ] ])".
    ParameterList   =   ParameterDecl   { "," ParameterDecl }
    ParameterDecl   =   [ IdentifierList ] [ "..."] Type.

- An interface type specifies a method set called its interface.
    InterfaceType   =   "interface" "{ { MethodSpec ";" }}"
    MethodSpec  =   MethodName  Signature | InterfaceTypeName.
    MethodName  =   identifier.
    InterfaceTypeName       =   TypeName.

- An interface type T may not embed itself or any interface type that embeds T, recursively.

- A map is an unordered group of elements of on type, called the element type,indexed by a set of unique keys of another type,called the key type.

    MapType =   "map" "[ KeyType ]" ElementType.
    KeyMap  =   Type

- The comparison operators == and != must be fully defined for operands of the key type.thus the key type must not be a function,map, or slice.if the key 
    type is the interface type,these comparison operators must be defined for the dynamic key values.failure will cause a run-time panic.

- map ( delete built-in function can remove element)

- a new empty value is made using the built-in function make.

- a nil map is equivalent to an empty map except that no emements may be added.

-  a channel provides a mechanism for concurrently executing functions to communicate by sending and receiving values of specified element type.

    ChannelType =   ( "chan" | "chan" "<-"| "<-" "chan") ElementType

- The optional <- operator specifies the cahnnel direction.a channel may be constrained only to send or only to receive by conversion or assignment.

- a new initialized channel value can be made using the built-in function make.
- if the capacity is zeor or absent, the channel is unbuffered and communication succeeds only when both a sender and receiver are ready. a nil channel is never
    ready for communication.

- a channel may be closed with close.
- the multi-valued assignment form of the receiver operator reports wheter a received value was send before the channel was closed.

- A single channel may be used in send statements,receive operations, and calls to the built-in functions cap and len by any number of goroutines without further synchronization.
- channels acs as first-in-first-out queues.
# unsafe
-  The built-in package unsafe,known to the compiler,provides facilities for low-level programming
	including operations that violate the type system. A pacakge using unsafe must be vetted 
	manually for type safety and may not be portable.

- built-in functions are predeclared.the built-in functions do not have standard Go types, so they 
	can only appear in call expressios;they cannot be used as function values.
-  for a channel c, the built-in function close(c) records that no more values will be sent on the
	channel.Tt is an error if c is a receive-only channel.Sending to or closing a closed channel	

### Handling panics
- Execution errors trigger a run-time panic equivalent to a call  of the built-in function panic.
- Two built-in functions panic and recover assist in reporting and handling run-time panics and program-defined error conditions.

- While executing  a function F, an explicit call to panic or a run-time panic terminates the execution of F. Any function deffered by F
    are then executed as usual. Next, any deffered functions run by F's caller are run, and so on to any deffered by the top-level function
    in the executing goroutine. At that point,the program is terminated and error condition is reported,including the value of the argument to panic.

### package
#net/http


### godoc

- 
