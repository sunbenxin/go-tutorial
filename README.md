# All
[go-toturial("https://github.com/sunbenxin/go-tutorial")]

[vim-toturial("https://github.com/sunbenxin/vim-tutorial")]

[pg-toturial("https://github.com/sunbenxin/postgres-tutorial")]

[tantan-toturial("https://git.oschina.net/sunbenxin/mytantan")]

[make-toturial("https://github.com/sunbenxin/make-tutorial")]

[bash-toturial("https://github.com/sunbenxin/bash-toturial")]

[dict-toturial("https://github.com/sunbenxin/dict-toturial")]

--------------------------------------------------------------------------------------------

# go-tutorial

## words to remember
 - qualified omitted explicit brace bracket encompasses extent
 - flirting culture co-founder Valentine

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
-  the blank identifier is represented by the underscore character _ .

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
#json
- json marshal must be export field can be encode

- The client must close response body when finished with it.
- For control HTTP client headers, redirect policy,and other settings,create a Client.

- For control over proxies,TLS configuration,keep-alives,compression,and other settings, create a Transport.
- Clients and Trasports are safe for concurrent use by multiple goroutines and for efficiency should only be created once and re-used.
- ListenAndServe starts an HTTP server with a given address and handler. The handler is usually nil, which means to use DefaultServeMux. Handle and HandleFunc add handlers to DefaultServeMux

- More cotrol over the server's behavior is available by creating a custom  Server.
- The http package has transparent support for the HTTP/2 protocol when using HTTPS.Programs that must disable HTTP/2 can do so by setting Transport.TLSNextProto (for clients) or Server.TLSNextProto (for servers) to a non-nil, empty map. Alternatively, the following GODEBUG environment variables are currently supported:
- if client.do return err == nil, need close resp.body self whatever read it or not.

# client
## time set
## proxy
## safe for concurrent use by multiple goroutines
## cookies
## redirects

# roundtripper
- an interface representing the ability to execute a single HTTP transaction, obtaining the Response for a given Request.
- must be safe for concurrent use by multiple goroutines.
- RoundTripper should not attempt to interpret the response.
- In particular, RountTripper must return err== nil if it obtained a repsonse, regardless of the repsonse's HTTP status code.
 A non-nil err should be reserved for failure to obtain a response. Similarly, RoundTrip should not attempt to handle higher-level protocol details such as redirects, authentication, or cookies.

- should not modify the request, except for consuming and closing the Request's Body.
- must always close the body, including on errors,but depending on the implementation may do so in a separate goroutine even after roundTrip returns.This means that callers wanting to
reuse the body by subsequent requests must arrange to wait for the close call before doing so.
- The request' url and header fields must be initialized.

# Transport is an implementaton of RoundTripper that supports HTTP. HTTPS, and HTTP proxies.
- by default, Transport caches connections for future re-use.This may leave many open connections when accessing many hosts.
This behavior can be managed using Transport's CloseIdleConnections method and the MaxIdleConnsPerHost and DisableKeepAlives fields.

- 

# time


### godoc

### statement
## swithch statement
- "Switch" statement provices multi-way execution. An expression or type specifier is compared to "cases" inside the switch to determine which branch to execute.

    SwitchStmt = ExprSwitchStmt | TypeSwitchStmt

- In an expression switch, the cases contain expressions  that are compared  against the value of the swithc expression. in a type switch,the cases contain types are
    compared against the type of a specially annotated switch expressioin. The switch expression is evaluated exactly once in a switch statement.

- In an expression switch,the switch expression are evaluated and the case expression, which need not be constant,are evaluated left-to-right and top-to-bottom; the first
  one that equals the switch expression triggers executioin of the statements of the associated case; the other cases are skipped.If no case matches and there is a "default" case
  its statement are executed.That can be at most one default case and it may appear anywhere in the "switch" statement.A missing switch expression is equivalent to the boolean value true;

    ExprSwitchStmt  = "switch" [ SimpleStmt ";" ] [ Expression ] "{" { ExprCaseClause }"}".
    ExprCaseClause =    ExprSwitchCase ":" StatementList
    ExprSwitchCase  =   "case" ExpressionList | "default"

- If the switch expression evaluated to an untyped constant, it is first converted to its default type; if it's an untyped boolean value, it is first converted to type bool.
  The predeclared untyped value nil cannot be used as a switch expression.

- If the case expressioin is untyped, it is first converted to the type of the switch expression.For each case expressioin x and the value t of the switch expression, x==t must be
  a valid comparison.

- A "fallthrough" statement may appear as the last statement of all but the last clause of an expression switch.

- The switch expression may be preceded by a simple statement,which executes before the expression is evaluated.

- A compiler may not allowed multiple case expressions evaluating to the same constant.

- A type switch is marked by a special  switch expression that has the form of a type assertion using the reserved word type rather than an actual type.
    eg: switch x.(type){
        //case
        }

- As with type assertion,  x must be of interface type,and each non-interface type T listed in a case must implement the type of x.The types listed in the cases of a type
  switch must all be different.

    TypeSwitchStmt      =   "switch" [ SimpleStmt ";" ] TypeSwitchGuard "{" { TypeCaseClause }"}".
    TypeSwitchGuard     =   [ identifier ":=" ] PrimaryExpr "." "(" "tyep" ")"
    TypeCaseClause      =  TypeSwitchCase ":" StatementList
    TypeSwitchCase      =   "case" TypeList | "default".
    TypeList            =   Type { "," Type }.

- The type TypeSwitchGuard may include a short variable declaration. When that form is used, the variable is declared at the beginning of the implicit block in each clause.
  In clauses with a case listing exactly one type, the variable has that type; otherwise the variablel has the type of the expression in the TypeSwitchGuard.

- The type in a case may be nil; that case is used when the expression in the TypeSwitchGuard is a nil interface value. There may be at most one nil case.

- The fallthrough is not permitted in a type switch.

### note
=======
- golang function does not support function reload(函数重载)
- string  literal represents a string constant obtained fro concatenating a sequence of characters.There are two forms:
	raw string literals and interpreted string literals.

- Raw literals are quoted by back quotes. Raw literals may contains newline.Carriage return character("\r") inside raw literals are discarded from then raw string value.
- Interpreted string are quoted by double quotes. newline can't inside interpreted strings.

- * string variable need to check if nil before dereference

### Coding style
- log
- if control

### Test



- clear commit message

- struct init // type1(type2)

- short variable declaration need has new variable on the left side.

# go in practice

## IDE
### bash
### vim

## fundamentals
### command-line

### handle configuration

### web server

### cache(redis,mongodb,memcache ...)

### database(postgres,mysql)
- pg client(gopkg.in/pg)
    - go-pg recognizes placeholders(`?`) in queries and replaces them with parameters when queries are executed.Parameters are escaped before replacing according to PostgreSQL rules.
        - all parameters are properly quoted against SQL injections
        - null byte is removed;
        - JSON/JSONB gets `\u0000` escaped as `\\u0000`.
    - support indexed parameters, named parameters,global params

### go concurrency
- two way: csp , shared variables

- Each concurrently executing activity is called goroutine.
- When a program starts,its only goroutine is the one that calls the main function,the main goroutine.
- There are no ways to communicate with a goroutine to request that it stop itself. And one goroutine to stop another.
- As with maps, a channel is a reference to the data structure created by make.When we copy a channel or pass one as an argument to a function,we are copying a reference.The zero value of a channel is nil.

- Two channels of the same type may be compared using ==.The comparision is true if both are references to the same channel data structure. A channel may also be compared to nil.
- A channel has two principal operations, send and receive. A receive expression whose result is not used is valid statement:

    ch <- x // a send statement
    x = <- ch // a receive expression in an assignment statement
    <-ch    // a receive statement ; result is discarded

- Channels support close operation.which sets a flag indicating that no more values will ever be sent on this channel,subsequent attempts to send will panic.
    Receive operations on a closed channel yield the values that have been sent until no more values are left; any receive operatons thereafter complete immediately and yield the zero value of the channel's element type.

- channel can be unbuffered channel and buffered channel.

    ch = make(chan int)         // unbuffered channel
    ch = make(chan int, 0)          // unbuffered channel
    ch = make(chan int, 3)          // buffered channel with capacity 3

- A send operation on an unbuffered channel blocks the sending goroutine until another goroutine executes a corresponding receive on the same channel, at which point the value is transmitted and both goroutines may continue.also the receive operation.

- Communication over an unbuffered channel causes the sending and receiving goroutines to synchronize.So unbuffered channels are sometimes called synchronous channels.

- There is no way to test directly whether a channel has been closed.but thereis a variant of the receive operation that produces two results: the received channel element, plus a bollean value.conventionally called ok.which is true for a successful receive and false for a receive on
    a closed and trained channel.conventionally called ok.which is true for a successful receive and false for a receive on
        a closed and drained channel.
- Need not to close channel, the garbage collector will do whether or not it is closed.
- Attempting to clos an already-closed channel causes a panic.as does close a nil channel. Close channels has another use as a broadcast mechanism.

- directional channel can limint its intent and prevent misuse.

- It is a mistake to use buffered channels within a single goroutine as a queue.Channels are deeply connected to goroutine scheduling,and without another goroutine receiving from the channel,a sender- and perhaps the whole program -risks becoming blocked forever.
- leak goroutines are not automatically collected.so it is important to make sure that goroutines terminate theselves when no longer needed.
- failure to allocate sufficient buffer capacity would cause the program to deadlock.

- When we konw an uppr bound on the number of values that will be sent on a channel,it is not unusual to create a buffered channel of that size and perform all the sends before the first value is received.

- Channel buffering may also affect program performance.
## mechanics of managing a Go application
### Go package management
- verdoring way
    - Go1.6 includes support for using local copies of external dependencies to satisfy imports of those dependencie, oftern referred to as vendoring.
    - Code below a directory named "vendor" is importable only by code in the directory tree rooted at the parent of "vendor",and only using an import path
    that omits the prefix up to and including the vendor element.
    - Code in vendor directories is not subject to import path checking.
    - When 'go get' checks out or updates a git repository, it now also updates submodules.

- version way by git repo
    - by directory (floder v1, v2...)

- gopkg.in supoort url redirect from a versioned url to code repo

### handling errors and panics
### debugging and testing(performance test)
- An error indicates that a particular task couldn’t be completed successfully
- A panic indicates that a severe event occurred, probably as a result of a programmer error

## user interfaces for you application
### html and email template patterns
### serving and receiving assests and forms

### REST APIs
### passing and handling errors over http
### passing and mapping JSON

### Versioning REST APIs

## talking your application to the cloud

### rpc


### redis
- concurrent i/o

### code review comments
https://github.com/golang/go/wiki/Comments

### code review
https://github.com/golang/go/wiki/CodeReviewComments


