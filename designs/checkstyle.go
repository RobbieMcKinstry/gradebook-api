package designs

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Checkstyle contains a full specification of which Checkstyle lints are to be checked.
var Checkstyle = Type("Checkstyle", func() {
	Description("Description of which Checkstyle lints are checked.")

	Attribute("AbbreviationAsWordInName", Boolean, "The Check validate abbreviations(consecutive capital letters) length in identifier name, it also allow in enforce camel case naming.")

	Attribute("AbstractClassName", Boolean, "Ensures that the names of abstract classes conforming to some regular expression.")

	Attribute("AnnotationLocation", Boolean, "Check location of annotation on language elements.")

	Attribute("AnnotationUseStyle", Boolean, "This check controls the style with the usage of annotations.")

	Attribute("AnonInnerLength", Boolean, "Checks for long anonymous inner classes.")

	Attribute("ArrayTrailingComma", Boolean, "Checks if array initialization contains optional trailing comma.")

	Attribute("ArrayTypeStyle", Boolean, "Checks the style of array type definitions.")

	Attribute("AtclauseOrder", Boolean, "Checks the order of at-clauses.")

	Attribute("AvoidEscapedUnicodeCharacters", Boolean, "Restrict using Unicode escapes.")

	Attribute("AvoidInlineConditionals", Boolean, "Detects inline conditionals.")

	Attribute("AvoidNestedBlocks", Boolean, "Finds nested blocks.")

	Attribute("AvoidStarImport", Boolean, "Check that finds import statements that use the * notation.")

	Attribute("AvoidStaticImport", Boolean, "Check that finds static imports.")

	Attribute("BooleanExpressionComplexity", Boolean, "Restricts nested boolean operators (&&, ||, &, | and ^) to a specified depth (default = 3).")

	Attribute("CatchParameterName", Boolean, "Checks that catch parameter names conform to a format specified by the format property.")

	Attribute("ClassDataAbstractionCoupling", Boolean, "This metric measures the number of instantiations of other classes within the given class.")

	Attribute("ClassFanOutComplexity", Boolean, "The number of other classes a given class relies on.")

	Attribute("ClassTypeParameterName", Boolean, "Checks that class type parameter names conform to a format specified by the format property.")

	Attribute("CommentsIndentation", Boolean, "Controls the indentation between comments and surrounding code.")

	Attribute("ConstantName", Boolean, "Checks that constant names conform to a format specified by the format property.")

	Attribute("CovariantEquals", Boolean, "Checks that if a class defines a covariant method equals, then it defines method equals(java.lang.Object).")

	Attribute("CustomImportOrder", Boolean, "Checks that the groups of import declarations appear in the order specified by the user.")

	Attribute("CyclomaticComplexity", Boolean, "Checks cyclomatic complexity against a specified limit.")

	Attribute("DeclarationOrder", Boolean, "Checks that the parts of a class or interface declaration appear in the order suggested by the Code Conventions for the Java Programming Language.")

	Attribute("DefaultComesLast", Boolean, "Check that the default is after all the cases in a switch statement.")

	Attribute("DescendantToken", Boolean, "Checks for restricted tokens beneath other tokens.")

	Attribute("DesignForExtension", Boolean, "Checks that classes are designed for inheritance.")

	Attribute("EmptyBlock", Boolean, "Checks for empty blocks.")

	Attribute("EmptyCatchBlock", Boolean, "Checks for empty catch blocks with few options to skip violation.")

	Attribute("EmptyForInitializerPad", Boolean, "Checks the padding of an empty for initializer; that is whether a space is required at an empty for initializer, or such spaces are forbidden.")

	Attribute("EmptyForIteratorPad", Boolean, "Checks the padding of an empty for iterator; that is whether a space is required at an empty for iterator, or such spaces are forbidden.")

	Attribute("EmptyLineSeparator", Boolean, "Checks for blank line separators.")

	Attribute("EmptyStatement", Boolean, "Detects empty statements (standalone ';').")

	Attribute("EqualsAvoidNull", Boolean, "Checks that any combination of String literals is on the left side of an equals() comparison.")

	Attribute("EqualsHashCode", Boolean, "Checks that classes that override equals() also override hashCode().")

	Attribute("ExecutableStatementCount", Boolean, "Restricts the number of executable statements to a specified limit (default = 30).")

	Attribute("ExplicitInitialization", Boolean, "Checks if any class or object member explicitly initialized to default for its type value (null for object references, zero for numeric types and char and false for boolean.")

	Attribute("FallThrough", Boolean, "Checks for fall through in switch statements Finds locations where a case contains Java code - but lacks a break, return, throw or continue statement.")

	Attribute("FileLength", Boolean, "Checks for long source files.")

	Attribute("FileTabCharacter", Boolean, "Checks to see if a file contains a tab character.")

	Attribute("FinalClass", Boolean, "Checks that class which has only private constructors is declared as final.")

	Attribute("FinalLocalVariable", Boolean, "Ensures that local variables that never get their values changed, must be declared final.")

	Attribute("FinalParameters", Boolean, "Check that method/constructor/catch/foreach parameters are final.")

	Attribute("GenericWhitespace", Boolean, "Checks that the whitespace around the Generic tokens < and > are correct to the typical convention.")

	Attribute("Header", Boolean, "Checks the header of the source against a fixed header file.")

	Attribute("HiddenField", Boolean, "Checks that a local variable or a parameter does not shadow a field that is defined in the same class.")

	Attribute("HideUtilityClassConstructor", Boolean, "Make sure that utility classes (classes that contain only static methods) do not have a public constructor.")

	Attribute("IllegalCatch", Boolean, "Catching java.lang.Exception, java.lang.Error or java.lang.RuntimeException is almost never acceptable.")

	Attribute("IllegalImport", Boolean, "Checks for imports from a set of illegal packages.")

	Attribute("IllegalInstantiation", Boolean, "Checks for illegal instantiations where a factory method is preferred.")

	Attribute("IllegalThrows", Boolean, "Throwing java.lang.Error or java.lang.RuntimeException is almost never acceptable.")

	Attribute("IllegalToken", Boolean, "Checks for illegal tokens.")

	Attribute("IllegalTokenText", Boolean, "Checks for illegal token text.")

	Attribute("IllegalType", Boolean, "Checks that particular class are never used as types in variable declarations, return values or parameters.")

	Attribute("ImportControl", Boolean, "Check that controls what packages can be imported in each package.")

	Attribute("ImportOrder", Boolean, "Ensures that groups of imports come in a specific order.")

	Attribute("Indentation", Boolean, "Checks correct indentation of Java Code.")

	Attribute("InnerAssignment", Boolean, "Checks for assignments in subexpressions, such as in String s = Integer.toString(i = 2);.")

	Attribute("InnerTypeLast", Boolean, "Check nested (internal) classes/interfaces are declared at the bottom of the class after all method and field declarations.")

	Attribute("InterfaceIsType", Boolean, "Implements Bloch, Effective Java, Item 17 - Use Interfaces only to define types.")

	Attribute("InterfaceTypeParameterName", Boolean, "Checks that interface type parameter names conform to a format specified by the format property.")

	Attribute("JavadocMethod", Boolean, "Checks the Javadoc of a method or constructor.")

	Attribute("JavadocPackage", Boolean, "Checks that all packages have a package documentation.")

	Attribute("JavadocParagraph", Boolean, "Checks Javadoc paragraphs.")

	Attribute("JavadocStyle", Boolean, "Custom Checkstyle Check to validate Javadoc.")

	Attribute("JavadocTagContinuationIndentation", Boolean, "Checks the indentation of the continuation lines in at-clauses.")

	Attribute("JavadocType", Boolean, "Checks the Javadoc of a type.")

	Attribute("JavadocVariable", Boolean, "Checks that a variable has Javadoc comment.")

	Attribute("JavaNCSS", Boolean, "This check calculates the Non Commenting Source Statements (NCSS) metric for Java source files and methods.")

	Attribute("LeftCurly", Boolean, "Checks the placement of left curly braces on types, methods and other blocks:")

	Attribute("LineLength", Boolean, "Checks for long lines.")

	Attribute("LocalFinalVariableName", Boolean, "Checks that local final variable names conform to a format specified by the format property.")

	Attribute("LocalVariableName", Boolean, "Checks that local, non-final variable names conform to a format specified by the format property.")

	Attribute("MagicNumber", Boolean, "Checks for magic numbers.")

	Attribute("MemberName", Boolean, "Checks that instance variable names conform to a format specified by the format property.")

	Attribute("MethodCount", Boolean, "Checks the number of methods declared in each type.")

	Attribute("MethodLength", Boolean, "Checks for long methods.")

	Attribute("MethodName", Boolean, "Checks that method names conform to a format specified by the format property.")

	Attribute("MethodParamPad", Boolean, "Checks the padding between the identifier of a method definition, constructor definition, method call, or constructor invocation; and the left parenthesis of the parameter list.")

	Attribute("MethodTypeParameterName", Boolean, "Checks that class type parameter names conform to a format specified by the format property.")

	Attribute("MissingCtor", Boolean, "Checks that classes (except abstract one) define a ctor and don't rely on the default one.")

	Attribute("MissingDeprecated", Boolean, "This class is used to verify that both the java.lang.Deprecated annotation is present and the @deprecated Javadoc tag is present when either is present.")

	Attribute("MissingOverride", Boolean, "This class is used to verify that the java.lang.Override annotation is present when the {@inheritDoc} javadoc tag is present.")

	Attribute("MissingSwitchDefault", Boolean, "Checks that switch statement has 'default' clause.")

	Attribute("ModifiedControlVariable", Boolean, "Check for ensuring that for loop control variables are not modified inside the for block.")

	Attribute("ModifierOrder", Boolean, "Checks that the order of modifiers conforms to the suggestions in the Java Language specification, sections 8.1.1, 8.3.1 and 8.4.3.")

	Attribute("MultipleStringLiterals", Boolean, "Checks for multiple occurrences of the same string literal within a single file.")

	Attribute("MultipleVariableDeclarations", Boolean, "Checks that each variable declaration is in its own statement and on its own line.")

	Attribute("MutableException", Boolean, "Ensures that exceptions (defined as any class name conforming to some regular expression) are immutable.")

	Attribute("NeedBraces", Boolean, "Checks for braces around code blocks.")

	Attribute("NestedForDepth", Boolean, "Restricts nested for blocks to a specified depth (default = 1).")

	Attribute("NestedIfDepth", Boolean, "Restricts nested if-else blocks to a specified depth (default = 1).")

	Attribute("NestedTryDepth", Boolean, "Restricts nested try-catch-finally blocks to a specified depth (default = 1).")

	Attribute("NewlineAtEndOfFile", Boolean, "Checks that there is a newline at the end of each file.")

	Attribute("NoClone", Boolean, "Checks that the clone method is not overridden from the Object class.")

	Attribute("NoFinalizer", Boolean, "Checks that no method having zero parameters is defined using the name finalize.")

	Attribute("NoLineWrap", Boolean, "Checks that chosen statements are not line-wrapped.")

	Attribute("NonEmptyAtclauseDescription", Boolean, "Checks that the at-clause tag is followed by description .")

	Attribute("NoWhitespaceAfter", Boolean, "Checks that there is no whitespace after a token.")

	Attribute("NoWhitespaceBefore", Boolean, "Checks that there is no whitespace before a token.")

	Attribute("NPathComplexity", Boolean, "Checks the npath complexity against a specified limit (default = 200).")

	Attribute("OneStatementPerLine", Boolean, "Checks that there is only one statement per line.")

	Attribute("OneTopLevelClass", Boolean, "Checks that each top-level class, interfaces or enum resides in a source file of its own.")

	Attribute("OperatorWrap", Boolean, "Checks line wrapping for operators.")

	Attribute("OuterTypeFilename", Boolean, "Checks that the outer type name and the file name match.")

	Attribute("OuterTypeNumber", Boolean, "Checks for the number of defined types at the 'outer' level.")

	Attribute("OverloadMethodsDeclarationOrder", Boolean, "Checks that overload methods are grouped together.")

	Attribute("PackageAnnotation", Boolean, "This check makes sure that all package annotations are in the package-info.java file.")

	Attribute("PackageDeclaration", Boolean, "Ensures there is a package declaration and (optionally) in the correct directory.")

	Attribute("PackageName", Boolean, "Checks that package names conform to a format specified by the format property.")

	Attribute("ParameterAssignment", Boolean, "Disallow assignment of parameters.")

	Attribute("ParameterName", Boolean, "Checks that parameter names conform to a format specified by the format property.")

	Attribute("ParameterNumber", Boolean, "Checks the number of parameters that a method or constructor has.")

	Attribute("ParenPad", Boolean, "Checks the padding of parentheses; that is whether a space is required after a left parenthesis and before a right parenthesis, or such spaces are forbidden, with the exception that it does not check for padding of the right parenthesis at an empty for iterator.")

	Attribute("RedundantImport", Boolean, "Checks for imports that are redundant.")

	Attribute("RedundantModifier", Boolean, "Checks for redundant modifiers in interface and annotation definitions, final modifier on methods of final classes, inner interface declarations that are declared as static, non public class constructors and enum constructors, nested enum definitions that are declared as static.")

	Attribute("Regexp", Boolean, "A check that makes sure that a specified pattern exists (or not) in the file.")

	Attribute("RegexpHeader", Boolean, "Checks the header of the source against a header file that contains a")

	Attribute("RegexpMultiline", Boolean, "Implementation of a check that looks that matches across multiple lines in any file type.")

	Attribute("RegexpOnFilename", Boolean, "Implementation of a check that matches based on file and/or folder path.")

	Attribute("RegexpSingleline", Boolean, "Implementation of a check that looks for a single line in any file type.")

	Attribute("RegexpSinglelineJava", Boolean, "Implementation of a check that looks for a single line in Java files.")

	Attribute("RequireThis", Boolean, "Checks that code doesn't rely on the 'this' default.")

	Attribute("ReturnCount", Boolean, "Restricts return statements to a specified count (default = 2).")

	Attribute("RightCurly", Boolean, "Checks the placement of right curly braces.")

	Attribute("SeparatorWrap", Boolean, "Checks line wrapping with separators.")

	Attribute("SimplifyBooleanExpression", Boolean, "Checks for overly complicated boolean expressions.")

	Attribute("SimplifyBooleanReturn", Boolean, "Checks for overly complicated boolean return statements.")

	Attribute("SingleLineJavadoc", Boolean, "Checks that a JavaDoc block which can fit on a single line and doesn't contain at-clauses")

	Attribute("SingleSpaceSeparator", Boolean, "Checks that non-whitespace characters are separated by no more than one whitespace.")

	Attribute("StaticVariableName", Boolean, "Checks that static, non-final variable names conform to a format specified by the format property.")

	Attribute("StringLiteralEquality", Boolean, "Checks that string literals are not used with == or !=.")

	Attribute("SummaryJavadoc", Boolean, "Checks that Javadoc summary sentence does not contain phrases that are not recommended to use.")

	Attribute("SuperClone", Boolean, "Checks that an overriding clone() method invokes super.clone().")

	Attribute("SuperFinalize", Boolean, "Checks that an overriding finalize() method invokes super.finalize().")

	Attribute("SuppressWarnings", Boolean, "This check allows you to specify what warnings that")

	Attribute("SuppressWarningsHolder", Boolean, "This check allows for finding code that should not be reported by Checkstyle")

	Attribute("ThrowsCount", Boolean, "Restricts throws statements to a specified count (default = 4).")

	Attribute("TodoComment", Boolean, "A check for TODO comments.")

	Attribute("TrailingComment", Boolean, "The check to ensure that requires that comments be the only thing on a line.")

	Attribute("Translation", Boolean, "The TranslationCheck class helps to ensure the correct translation of code by checking property files for consistency regarding their keys.")

	Attribute("TypecastParenPad", Boolean, "Checks the padding of parentheses for typecasts.")

	Attribute("TypeName", Boolean, "Checks that type names conform to a format specified by the format property.")

	Attribute("UncommentedMain", Boolean, "Detects uncommented main methods.")

	Attribute("UniqueProperties", Boolean, "Detects duplicated keys in properties files.")

	Attribute("UnnecessaryParentheses", Boolean, "Checks if unnecessary parentheses are used in a statement or expression.")

	Attribute("UnusedImports", Boolean, "Checks for unused import statements.")

	Attribute("UpperEll", Boolean, "Checks that long constants are defined with an upper ell.")

	Attribute("VariableDeclarationUsageDistance", Boolean, "Checks the distance between declaration of variable and its first usage.")

	Attribute("VisibilityModifier", Boolean, "Checks visibility of class members.")

	Attribute("WhitespaceAfter", Boolean, "Checks that a token is followed by whitespace, with the exception that it does not check for whitespace after the semicolon of an empty for iterator.")

	Attribute("WhitespaceAround", Boolean, "Checks that a token is surrounded by whitespace.")

	Attribute("WriteTag", Boolean, "Outputs a JavaDoc tag as information.")
})
