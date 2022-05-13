package ast

type SyntaxKind struct {
	FirstToken                                   int
	EndOfFileToken                               int
	FirstTriviaToken                             int
	MultiLineCommentTrivia                       int
	NewLineTrivia                                int
	WhitespaceTrivia                             int
	ShebangTrivia                                int
	LastTriviaToken                              int
	FirstLiteralToken                            int
	BigIntLiteral                                int
	intLiteral                                   int
	JsxText                                      int
	JsxTextAllWhiteSpaces                        int
	RegularExpressionLiteral                     int
	FirstTemplateToken                           int
	TemplateHead                                 int
	TemplateMiddle                               int
	LastTemplateToken                            int
	FirstPunctuation                             int
	CloseBraceToken                              int
	OpenParenToken                               int
	CloseParenToken                              int
	OpenBracketToken                             int
	CloseBracketToken                            int
	DotToken                                     int
	DotDotDotToken                               int
	SemicolonToken                               int
	CommaToken                                   int
	QuestionDotToken                             int
	FirstBinaryOperator                          int
	LessThanSlashToken                           int
	GreaterThanToken                             int
	LessThanEqualsToken                          int
	GreaterThanEqualsToken                       int
	EqualsEqualsToken                            int
	ExclamationEqualsToken                       int
	EqualsEqualsEqualsToken                      int
	ExclamationEqualsEqualsToken                 int
	EqualsGreaterThanToken                       int
	PlusToken                                    int
	MinusToken                                   int
	AsteriskToken                                int
	AsteriskAsteriskToken                        int
	SlashToken                                   int
	PercentToken                                 int
	PlusPlusToken                                int
	MinusMinusToken                              int
	LessThanLessThanToken                        int
	GreaterThanGreaterThanToken                  int
	GreaterThanGreaterThanGreaterThanToken       int
	AmpersandToken                               int
	BarToken                                     int
	CaretToken                                   int
	ExclamationToken                             int
	TildeToken                                   int
	AmpersandAmpersandToken                      int
	BarBarToken                                  int
	QuestionToken                                int
	ColonToken                                   int
	AtToken                                      int
	QuestionQuestionToken                        int
	BacktickToken                                int
	HashToken                                    int
	FirstAssignment                              int
	FirstCompoundAssignment                      int
	MinusEqualsToken                             int
	AsteriskEqualsToken                          int
	AsteriskAsteriskEqualsToken                  int
	SlashEqualsToken                             int
	PercentEqualsToken                           int
	LessThanLessThanEqualsToken                  int
	GreaterThanGreaterThanEqualsToken            int
	GreaterThanGreaterThanGreaterThanEqualsToken int
	AmpersandEqualsToken                         int
	BarEqualsToken                               int
	BarBarEqualsToken                            int
	AmpersandAmpersandEqualsToken                int
	QuestionQuestionEqualsToken                  int
	LastBinaryOperator                           int
	Identifier                                   int
	PrivateIdentifier                            int
	FirstKeyword                                 int
	CaseKeyword                                  int
	CatchKeyword                                 int
	ClassKeyword                                 int
	ConstKeyword                                 int
	ContinueKeyword                              int
	DebuggerKeyword                              int
	DefaultKeyword                               int
	DeleteKeyword                                int
	DoKeyword                                    int
	ElseKeyword                                  int
	EnumKeyword                                  int
	ExportKeyword                                int
	ExtendsKeyword                               int
	FalseKeyword                                 int
	FinallyKeyword                               int
	ForKeyword                                   int
	FunctionKeyword                              int
	IfKeyword                                    int
	ImportKeyword                                int
	InKeyword                                    int
	InstanceOfKeyword                            int
	NewKeyword                                   int
	NullKeyword                                  int
	ReturnKeyword                                int
	SuperKeyword                                 int
	SwitchKeyword                                int
	ThisKeyword                                  int
	ThrowKeyword                                 int
	TrueKeyword                                  int
	TryKeyword                                   int
	TypeOfKeyword                                int
	VarKeyword                                   int
	VoidKeyword                                  int
	WhileKeyword                                 int
	LastReservedWord                             int
	FirstFutureReservedWord                      int
	InterfaceKeyword                             int
	LetKeyword                                   int
	PackageKeyword                               int
	PrivateKeyword                               int
	ProtectedKeyword                             int
	PublicKeyword                                int
	StaticKeyword                                int
	LastFutureReservedWord                       int
	FirstContextualKeyword                       int
	AsKeyword                                    int
	AssertsKeyword                               int
	AssertKeyword                                int
	AnyKeyword                                   int
	AsyncKeyword                                 int
	AwaitKeyword                                 int
	BooleanKeyword                               int
	ConstructorKeyword                           int
	DeclareKeyword                               int
	GetKeyword                                   int
	InferKeyword                                 int
	IntrinsicKeyword                             int
	IsKeyword                                    int
	KeyOfKeyword                                 int
	ModuleKeyword                                int
	NamespaceKeyword                             int
	NeverKeyword                                 int
	ReadonlyKeyword                              int
	RequireKeyword                               int
	NumberKeyword                                int
	ObjectKeyword                                int
	SetKeyword                                   int
	intKeyword                                   int
	SymbolKeyword                                int
	TypeKeyword                                  int
	UndefinedKeyword                             int
	UniqueKeyword                                int
	UnknownKeyword                               int
	FromKeyword                                  int
	GlobalKeyword                                int
	BigIntKeyword                                int
	OverrideKeyword                              int
	LastContextualKeyword                        int
	FirstNode                                    int
	ComputedPropertyName                         int
	TypeParameter                                int
	Parameter                                    int
	Decorator                                    int
	PropertySignature                            int
	PropertyDeclaration                          int
	MethodSignature                              int
	MethodDeclaration                            int
	ClassStaticBlockDeclaration                  int
	Constructor                                  int
	GetAccessor                                  int
	SetAccessor                                  int
	CallSignature                                int
	ConstructSignature                           int
	IndexSignature                               int
	FirstTypeNode                                int
	TypeReference                                int
	FunctionType                                 int
	ConstructorType                              int
	TypeQuery                                    int
	TypeLiteral                                  int
	ArrayType                                    int
	TupleType                                    int
	OptionalType                                 int
	RestType                                     int
	UnionType                                    int
	IntersectionType                             int
	ConditionalType                              int
	InferType                                    int
	ParenthesizedType                            int
	ThisType                                     int
	TypeOperator                                 int
	IndexedAccessType                            int
	MappedType                                   int
	LiteralType                                  int
	NamedTupleMember                             int
	TemplateLiteralType                          int
	TemplateLiteralTypeSpan                      int
	LastTypeNode                                 int
	ObjectBindingPattern                         int
	ArrayBindingPattern                          int
	BindingElement                               int
	ArrayLiteralExpression                       int
	ObjectLiteralExpression                      int
	PropertyAccessExpression                     int
	ElementAccessExpression                      int
	CallExpression                               int
	NewExpression                                int
	TaggedTemplateExpression                     int
	TypeAssertionExpression                      int
	ParenthesizedExpression                      int
	FunctionExpression                           int
	ArrowFunction                                int
	DeleteExpression                             int
	TypeOfExpression                             int
	VoidExpression                               int
	AwaitExpression                              int
	PrefixUnaryExpression                        int
	PostfixUnaryExpression                       int
	BinaryExpression                             int
	ConditionalExpression                        int
	TemplateExpression                           int
	YieldExpression                              int
	SpreadElement                                int
	ClassExpression                              int
	OmittedExpression                            int
	ExpressionWithTypeArguments                  int
	AsExpression                                 int
	NonNullExpression                            int
	MetaProperty                                 int
	SyntheticExpression                          int
	TemplateSpan                                 int
	SemicolonClassElement                        int
	Block                                        int
	EmptyStatement                               int
	FirstStatement                               int
	ExpressionStatement                          int
	IfStatement                                  int
	DoStatement                                  int
	WhileStatement                               int
	ForStatement                                 int
	ForInStatement                               int
	ForOfStatement                               int
	ContinueStatement                            int
	BreakStatement                               int
	ReturnStatement                              int
	WithStatement                                int
	SwitchStatement                              int
	LabeledStatement                             int
	ThrowStatement                               int
	TryStatement                                 int
	LastStatement                                int
	VariableDeclaration                          int
	VariableDeclarationList                      int
	FunctionDeclaration                          int
	ClassDeclaration                             int
	InterfaceDeclaration                         int
	TypeAliasDeclaration                         int
	EnumDeclaration                              int
	ModuleDeclaration                            int
	ModuleBlock                                  int
	CaseBlock                                    int
	NamespaceExportDeclaration                   int
	ImportEqualsDeclaration                      int
	ImportDeclaration                            int
	ImportClause                                 int
	NamespaceImport                              int
	NamedImports                                 int
	ImportSpecifier                              int
	ExportAssignment                             int
	ExportDeclaration                            int
	NamedExports                                 int
	NamespaceExport                              int
	ExportSpecifier                              int
	MissingDeclaration                           int
	ExternalModuleReference                      int
	JsxElement                                   int
	JsxSelfClosingElement                        int
	JsxOpeningElement                            int
	JsxClosingElement                            int
	JsxFragment                                  int
	JsxOpeningFragment                           int
	JsxClosingFragment                           int
	JsxAttribute                                 int
	JsxAttributes                                int
	JsxSpreadAttribute                           int
	JsxExpression                                int
	CaseClause                                   int
	DefaultClause                                int
	HeritageClause                               int
	CatchClause                                  int
	AssertClause                                 int
	AssertEntry                                  int
	PropertyAssignment                           int
	ShorthandPropertyAssignment                  int
	SpreadAssignment                             int
	EnumMember                                   int
	UnparsedPrologue                             int
	UnparsedPrepend                              int
	UnparsedText                                 int
	UnparsedInternalText                         int
	UnparsedSyntheticReference                   int
	SourceFile                                   int
	Bundle                                       int
	UnparsedSource                               int
	InputFiles                                   int
	FirstJSDocNode                               int
	JSDocNameReference                           int
	JSDocMemberName                              int
	JSDocAllType                                 int
	JSDocUnknownType                             int
	JSDocNullableType                            int
	JSDocNonNullableType                         int
	JSDocOptionalType                            int
	JSDocFunctionType                            int
	JSDocVariadicType                            int
	JSDocNamepathType                            int
	JSDocComment                                 int
	JSDocText                                    int
	JSDocTypeLiteral                             int
	JSDocSignature                               int
	JSDocLink                                    int
	JSDocLinkCode                                int
	JSDocLinkPlain                               int
	FirstJSDocTagNode                            int
	JSDocAugmentsTag                             int
	JSDocImplementsTag                           int
	JSDocAuthorTag                               int
	JSDocDeprecatedTag                           int
	JSDocClassTag                                int
	JSDocPublicTag                               int
	JSDocPrivateTag                              int
	JSDocProtectedTag                            int
	JSDocReadonlyTag                             int
	JSDocOverrideTag                             int
	JSDocCallbackTag                             int
	JSDocEnumTag                                 int
	JSDocParameterTag                            int
	JSDocReturnTag                               int
	JSDocThisTag                                 int
	JSDocTypeTag                                 int
	JSDocTemplateTag                             int
	JSDocTypedefTag                              int
	JSDocSeeTag                                  int
	LastJSDocTagNode                             int
	SyntaxList                                   int
	NotEmittedStatement                          int
	PartiallyEmittedExpression                   int
	CommaListExpression                          int
	MergeDeclarationMarker                       int
	EndOfDeclarationMarker                       int
	SyntheticReferenceExpression                 int
	Count                                        int
}