// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/alligrader/gradebook-api/designs
// --out=$(GOPATH)/src/github.com/alligrader/gradebook-api/src
// --version=v1.1.0-dirty
//
// API "GradebookAPI": CLI Commands
//
// The content of this file is auto-generated, DO NOT MODIFY

package cli

import (
	"encoding/json"
	"fmt"
	"github.com/alligrader/gradebook-api/src/client"
	"github.com/goadesign/goa"
	goaclient "github.com/goadesign/goa/client"
	uuid "github.com/goadesign/goa/uuid"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type (
	// CreateBugProfileCommand is the command line data structure for the create action of bug_profile
	CreateBugProfileCommand struct {
		Payload     string
		ContentType string
		PrettyPrint bool
	}

	// DeleteBugProfileCommand is the command line data structure for the delete action of bug_profile
	DeleteBugProfileCommand struct {
		// The profile's unique identifier
		ProfileID   int
		PrettyPrint bool
	}

	// ListBugProfileCommand is the command line data structure for the list action of bug_profile
	ListBugProfileCommand struct {
		PrettyPrint bool
	}

	// ShowBugProfileCommand is the command line data structure for the show action of bug_profile
	ShowBugProfileCommand struct {
		// The profile's unique identifier
		ProfileID   int
		PrettyPrint bool
	}

	// UpdateBugProfileCommand is the command line data structure for the update action of bug_profile
	UpdateBugProfileCommand struct {
		Payload     string
		ContentType string
		// The profile's unique identifier
		ProfileID   int
		PrettyPrint bool
	}

	// CreateUserCommand is the command line data structure for the create action of user
	CreateUserCommand struct {
		Payload     string
		ContentType string
		PrettyPrint bool
	}
)

// RegisterCommands registers the resource action CLI commands.
func RegisterCommands(app *cobra.Command, c *client.Client) {
	var command, sub *cobra.Command
	command = &cobra.Command{
		Use:   "create",
		Short: `create action`,
	}
	tmp1 := new(CreateBugProfileCommand)
	sub = &cobra.Command{
		Use:   `bug_profile ["/api/bug-profile"]`,
		Short: `Bug profile represents the list of bugs the user wants us to check for`,
		Long: `Bug profile represents the list of bugs the user wants us to check for

Payload example:

{
   "checkstyle": {
      "AbbreviationAsWordInName": true,
      "AbstractClassName": true,
      "AnnotationLocation": false,
      "AnnotationUseStyle": true,
      "AnonInnerLength": true,
      "ArrayTrailingComma": true,
      "ArrayTypeStyle": true,
      "AtclauseOrder": true,
      "AvoidEscapedUnicodeCharacters": true,
      "AvoidInlineConditionals": false,
      "AvoidNestedBlocks": true,
      "AvoidStarImport": true,
      "AvoidStaticImport": false,
      "BooleanExpressionComplexity": false,
      "CatchParameterName": true,
      "ClassDataAbstractionCoupling": false,
      "ClassFanOutComplexity": false,
      "ClassTypeParameterName": true,
      "CommentsIndentation": true,
      "ConstantName": false,
      "CovariantEquals": false,
      "CustomImportOrder": true,
      "CyclomaticComplexity": true,
      "DeclarationOrder": true,
      "DefaultComesLast": false,
      "DescendantToken": false,
      "DesignForExtension": true,
      "EmptyBlock": true,
      "EmptyCatchBlock": true,
      "EmptyForInitializerPad": false,
      "EmptyForIteratorPad": false,
      "EmptyLineSeparator": true,
      "EmptyStatement": false,
      "EqualsAvoidNull": true,
      "EqualsHashCode": true,
      "ExecutableStatementCount": true,
      "ExplicitInitialization": false,
      "FallThrough": false,
      "FileLength": false,
      "FileTabCharacter": false,
      "FinalClass": true,
      "FinalLocalVariable": false,
      "FinalParameters": false,
      "GenericWhitespace": true,
      "Header": false,
      "HiddenField": true,
      "HideUtilityClassConstructor": true,
      "IllegalCatch": false,
      "IllegalImport": true,
      "IllegalInstantiation": true,
      "IllegalThrows": true,
      "IllegalToken": true,
      "IllegalTokenText": false,
      "IllegalType": true,
      "ImportControl": true,
      "ImportOrder": true,
      "Indentation": false,
      "InnerAssignment": false,
      "InnerTypeLast": true,
      "InterfaceIsType": true,
      "InterfaceTypeParameterName": false,
      "JavaNCSS": false,
      "JavadocMethod": false,
      "JavadocPackage": true,
      "JavadocParagraph": true,
      "JavadocStyle": true,
      "JavadocTagContinuationIndentation": true,
      "JavadocType": true,
      "JavadocVariable": false,
      "LeftCurly": false,
      "LineLength": true,
      "LocalFinalVariableName": false,
      "LocalVariableName": true,
      "MagicNumber": true,
      "MemberName": false,
      "MethodCount": false,
      "MethodLength": true,
      "MethodName": false,
      "MethodParamPad": false,
      "MethodTypeParameterName": false,
      "MissingCtor": false,
      "MissingDeprecated": true,
      "MissingOverride": true,
      "MissingSwitchDefault": false,
      "ModifiedControlVariable": false,
      "ModifierOrder": false,
      "MultipleStringLiterals": false,
      "MultipleVariableDeclarations": false,
      "MutableException": true,
      "NPathComplexity": true,
      "NeedBraces": true,
      "NestedForDepth": false,
      "NestedIfDepth": true,
      "NestedTryDepth": false,
      "NewlineAtEndOfFile": true,
      "NoClone": false,
      "NoFinalizer": true,
      "NoLineWrap": true,
      "NoWhitespaceAfter": true,
      "NoWhitespaceBefore": true,
      "NonEmptyAtclauseDescription": true,
      "OneStatementPerLine": false,
      "OneTopLevelClass": false,
      "OperatorWrap": false,
      "OuterTypeFilename": false,
      "OuterTypeNumber": false,
      "OverloadMethodsDeclarationOrder": false,
      "PackageAnnotation": true,
      "PackageDeclaration": true,
      "PackageName": true,
      "ParameterAssignment": true,
      "ParameterName": false,
      "ParameterNumber": true,
      "ParenPad": false,
      "RedundantImport": false,
      "RedundantModifier": true,
      "Regexp": false,
      "RegexpHeader": true,
      "RegexpMultiline": false,
      "RegexpOnFilename": false,
      "RegexpSingleline": false,
      "RegexpSinglelineJava": true,
      "RequireThis": true,
      "ReturnCount": true,
      "RightCurly": true,
      "SeparatorWrap": false,
      "SimplifyBooleanExpression": false,
      "SimplifyBooleanReturn": true,
      "SingleLineJavadoc": true,
      "SingleSpaceSeparator": true,
      "StaticVariableName": true,
      "StringLiteralEquality": false,
      "SummaryJavadoc": false,
      "SuperClone": false,
      "SuperFinalize": false,
      "SuppressWarnings": true,
      "SuppressWarningsHolder": false,
      "ThrowsCount": false,
      "TodoComment": false,
      "TrailingComment": true,
      "Translation": true,
      "TypeName": true,
      "TypecastParenPad": false,
      "UncommentedMain": false,
      "UniqueProperties": true,
      "UnnecessaryParentheses": true,
      "UnusedImports": true,
      "UpperEll": true,
      "VariableDeclarationUsageDistance": true,
      "VisibilityModifier": false,
      "WhitespaceAfter": true,
      "WhitespaceAround": true,
      "WriteTag": true
   },
   "findbugs": {
      "name": "Dolores et aspernatur eos adipisci repellendus."
   },
   "name": "Aut architecto maxime dignissimos."
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp1.Run(c, args) },
	}
	tmp1.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp1.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp2 := new(CreateUserCommand)
	sub = &cobra.Command{
		Use:   `user ["/api/user"]`,
		Short: `Represents a user of the site`,
		Long: `Represents a user of the site

Payload example:

{
   "email": "francesca_hermiston@kilback.biz",
   "name": "Et minus odio.",
   "password": "Neque suscipit natus et eligendi alias.",
   "phone_number": "Molestias aut consequatur vitae porro."
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp2.Run(c, args) },
	}
	tmp2.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp2.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "delete",
		Short: `I don't want this bug profile anymore`,
	}
	tmp3 := new(DeleteBugProfileCommand)
	sub = &cobra.Command{
		Use:   `bug_profile ["/api/bug-profile/PROFILEID"]`,
		Short: `Bug profile represents the list of bugs the user wants us to check for`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp3.Run(c, args) },
	}
	tmp3.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp3.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "list",
		Short: `Show all of my bug profiles`,
	}
	tmp4 := new(ListBugProfileCommand)
	sub = &cobra.Command{
		Use:   `bug_profile ["/api/bug-profile"]`,
		Short: `Bug profile represents the list of bugs the user wants us to check for`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp4.Run(c, args) },
	}
	tmp4.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp4.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "show",
		Short: `Show a single bug profile`,
	}
	tmp5 := new(ShowBugProfileCommand)
	sub = &cobra.Command{
		Use:   `bug_profile ["/api/bug-profile/PROFILEID"]`,
		Short: `Bug profile represents the list of bugs the user wants us to check for`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp5.Run(c, args) },
	}
	tmp5.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp5.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "update",
		Short: `Update my pre-existing bug profile`,
	}
	tmp6 := new(UpdateBugProfileCommand)
	sub = &cobra.Command{
		Use:   `bug_profile ["/api/bug-profile/PROFILEID"]`,
		Short: `Bug profile represents the list of bugs the user wants us to check for`,
		Long: `Bug profile represents the list of bugs the user wants us to check for

Payload example:

{
   "checkstyle": {
      "AbbreviationAsWordInName": true,
      "AbstractClassName": true,
      "AnnotationLocation": false,
      "AnnotationUseStyle": true,
      "AnonInnerLength": true,
      "ArrayTrailingComma": true,
      "ArrayTypeStyle": true,
      "AtclauseOrder": true,
      "AvoidEscapedUnicodeCharacters": true,
      "AvoidInlineConditionals": false,
      "AvoidNestedBlocks": true,
      "AvoidStarImport": true,
      "AvoidStaticImport": false,
      "BooleanExpressionComplexity": false,
      "CatchParameterName": true,
      "ClassDataAbstractionCoupling": false,
      "ClassFanOutComplexity": false,
      "ClassTypeParameterName": true,
      "CommentsIndentation": true,
      "ConstantName": false,
      "CovariantEquals": false,
      "CustomImportOrder": true,
      "CyclomaticComplexity": true,
      "DeclarationOrder": true,
      "DefaultComesLast": false,
      "DescendantToken": false,
      "DesignForExtension": true,
      "EmptyBlock": true,
      "EmptyCatchBlock": true,
      "EmptyForInitializerPad": false,
      "EmptyForIteratorPad": false,
      "EmptyLineSeparator": true,
      "EmptyStatement": false,
      "EqualsAvoidNull": true,
      "EqualsHashCode": true,
      "ExecutableStatementCount": true,
      "ExplicitInitialization": false,
      "FallThrough": false,
      "FileLength": false,
      "FileTabCharacter": false,
      "FinalClass": true,
      "FinalLocalVariable": false,
      "FinalParameters": false,
      "GenericWhitespace": true,
      "Header": false,
      "HiddenField": true,
      "HideUtilityClassConstructor": true,
      "IllegalCatch": false,
      "IllegalImport": true,
      "IllegalInstantiation": true,
      "IllegalThrows": true,
      "IllegalToken": true,
      "IllegalTokenText": false,
      "IllegalType": true,
      "ImportControl": true,
      "ImportOrder": true,
      "Indentation": false,
      "InnerAssignment": false,
      "InnerTypeLast": true,
      "InterfaceIsType": true,
      "InterfaceTypeParameterName": false,
      "JavaNCSS": false,
      "JavadocMethod": false,
      "JavadocPackage": true,
      "JavadocParagraph": true,
      "JavadocStyle": true,
      "JavadocTagContinuationIndentation": true,
      "JavadocType": true,
      "JavadocVariable": false,
      "LeftCurly": false,
      "LineLength": true,
      "LocalFinalVariableName": false,
      "LocalVariableName": true,
      "MagicNumber": true,
      "MemberName": false,
      "MethodCount": false,
      "MethodLength": true,
      "MethodName": false,
      "MethodParamPad": false,
      "MethodTypeParameterName": false,
      "MissingCtor": false,
      "MissingDeprecated": true,
      "MissingOverride": true,
      "MissingSwitchDefault": false,
      "ModifiedControlVariable": false,
      "ModifierOrder": false,
      "MultipleStringLiterals": false,
      "MultipleVariableDeclarations": false,
      "MutableException": true,
      "NPathComplexity": true,
      "NeedBraces": true,
      "NestedForDepth": false,
      "NestedIfDepth": true,
      "NestedTryDepth": false,
      "NewlineAtEndOfFile": true,
      "NoClone": false,
      "NoFinalizer": true,
      "NoLineWrap": true,
      "NoWhitespaceAfter": true,
      "NoWhitespaceBefore": true,
      "NonEmptyAtclauseDescription": true,
      "OneStatementPerLine": false,
      "OneTopLevelClass": false,
      "OperatorWrap": false,
      "OuterTypeFilename": false,
      "OuterTypeNumber": false,
      "OverloadMethodsDeclarationOrder": false,
      "PackageAnnotation": true,
      "PackageDeclaration": true,
      "PackageName": true,
      "ParameterAssignment": true,
      "ParameterName": false,
      "ParameterNumber": true,
      "ParenPad": false,
      "RedundantImport": false,
      "RedundantModifier": true,
      "Regexp": false,
      "RegexpHeader": true,
      "RegexpMultiline": false,
      "RegexpOnFilename": false,
      "RegexpSingleline": false,
      "RegexpSinglelineJava": true,
      "RequireThis": true,
      "ReturnCount": true,
      "RightCurly": true,
      "SeparatorWrap": false,
      "SimplifyBooleanExpression": false,
      "SimplifyBooleanReturn": true,
      "SingleLineJavadoc": true,
      "SingleSpaceSeparator": true,
      "StaticVariableName": true,
      "StringLiteralEquality": false,
      "SummaryJavadoc": false,
      "SuperClone": false,
      "SuperFinalize": false,
      "SuppressWarnings": true,
      "SuppressWarningsHolder": false,
      "ThrowsCount": false,
      "TodoComment": false,
      "TrailingComment": true,
      "Translation": true,
      "TypeName": true,
      "TypecastParenPad": false,
      "UncommentedMain": false,
      "UniqueProperties": true,
      "UnnecessaryParentheses": true,
      "UnusedImports": true,
      "UpperEll": true,
      "VariableDeclarationUsageDistance": true,
      "VisibilityModifier": false,
      "WhitespaceAfter": true,
      "WhitespaceAround": true,
      "WriteTag": true
   },
   "findbugs": {
      "name": "Dolores et aspernatur eos adipisci repellendus."
   },
   "name": "Aut architecto maxime dignissimos."
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp6.Run(c, args) },
	}
	tmp6.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp6.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
}

func intFlagVal(name string, parsed int) *int {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func float64FlagVal(name string, parsed float64) *float64 {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func boolFlagVal(name string, parsed bool) *bool {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func stringFlagVal(name string, parsed string) *string {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func hasFlag(name string) bool {
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "--"+name) {
			return true
		}
	}
	return false
}

func jsonVal(val string) (*interface{}, error) {
	var t interface{}
	err := json.Unmarshal([]byte(val), &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func jsonArray(ins []string) ([]interface{}, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []interface{}
	for _, id := range ins {
		val, err := jsonVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, val)
	}
	return vals, nil
}

func timeVal(val string) (*time.Time, error) {
	t, err := time.Parse(time.RFC3339, val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func timeArray(ins []string) ([]time.Time, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []time.Time
	for _, id := range ins {
		val, err := timeVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func uuidVal(val string) (*uuid.UUID, error) {
	t, err := uuid.FromString(val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func uuidArray(ins []string) ([]uuid.UUID, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []uuid.UUID
	for _, id := range ins {
		val, err := uuidVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func float64Val(val string) (*float64, error) {
	t, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func float64Array(ins []string) ([]float64, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []float64
	for _, id := range ins {
		val, err := float64Val(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func boolVal(val string) (*bool, error) {
	t, err := strconv.ParseBool(val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func boolArray(ins []string) ([]bool, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []bool
	for _, id := range ins {
		val, err := boolVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

// Run makes the HTTP request corresponding to the CreateBugProfileCommand command.
func (cmd *CreateBugProfileCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/api/bug-profile"
	}
	var payload client.BugProfile
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.CreateBugProfile(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateBugProfileCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
}

// Run makes the HTTP request corresponding to the DeleteBugProfileCommand command.
func (cmd *DeleteBugProfileCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/bug-profile/%v", cmd.ProfileID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.DeleteBugProfile(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *DeleteBugProfileCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var profileID int
	cc.Flags().IntVar(&cmd.ProfileID, "profileID", profileID, `The profile's unique identifier`)
}

// Run makes the HTTP request corresponding to the ListBugProfileCommand command.
func (cmd *ListBugProfileCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/api/bug-profile"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ListBugProfile(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListBugProfileCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the ShowBugProfileCommand command.
func (cmd *ShowBugProfileCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/bug-profile/%v", cmd.ProfileID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ShowBugProfile(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowBugProfileCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var profileID int
	cc.Flags().IntVar(&cmd.ProfileID, "profileID", profileID, `The profile's unique identifier`)
}

// Run makes the HTTP request corresponding to the UpdateBugProfileCommand command.
func (cmd *UpdateBugProfileCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/bug-profile/%v", cmd.ProfileID)
	}
	var payload client.BugProfile
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.UpdateBugProfile(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UpdateBugProfileCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
	var profileID int
	cc.Flags().IntVar(&cmd.ProfileID, "profileID", profileID, `The profile's unique identifier`)
}

// Run makes the HTTP request corresponding to the CreateUserCommand command.
func (cmd *CreateUserCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/api/user"
	}
	var payload client.User
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.CreateUser(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateUserCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
}