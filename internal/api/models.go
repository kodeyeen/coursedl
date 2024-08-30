package api

import (
	"time"
)

type Course struct {
	ID                  string   `json:"_id"`
	Title               string   `json:"title"`
	Description         string   `json:"description"`
	CoverImageURL       string   `json:"coverImageUrl"`
	OriginalPrice       int      `json:"originalPrice"`
	DiscountPercent     int      `json:"discountPercent"`
	Validity            string   `json:"validity"`
	Type                string   `json:"type"`
	Courses             []string `json:"courses"`
	SortOrder           int      `json:"sortOrder"`
	Slug                string   `json:"slug"`
	Category            []any    `json:"category"`
	CountryDiscountType string   `json:"countryDiscountType,omitempty"`
	IsNewArrival        bool     `json:"isNewArrival"`
}

type Lesson struct {
	Title        string            `json:"title"`
	Created      time.Time         `json:"created"`
	Published    bool              `json:"published"`
	IsNewArrival bool              `json:"isNewArrival"`
	Documents    []*LessonDocument `json:"documents"`
	ID           string            `json:"_id"`
}

type LessonDocument struct {
	DocumentTitle         string `json:"documentTitle"`
	DocumentID            string `json:"documentId"`
	Published             bool   `json:"published"`
	IsPublic              bool   `json:"isPublic"`
	Slug                  string `json:"slug"`
	DiscussionID          string `json:"discussionId"`
	ReadOnly              bool   `json:"readOnly"`
	SupportedLanguages    []any  `json:"supportedLanguages"`
	DocumentTitleEmphasis bool   `json:"documentTitleEmphasis"`
	IsBlurredDocument     bool   `json:"isBlurredDocument"`
	IsNewArrivalDocument  bool   `json:"isNewArrivalDocument"`
	InternalID            string `json:"_id"`
	ID                    string `json:"id"`
}

type DocumentBody struct {
	LatestDocument struct {
		Updated  time.Time          `json:"updated"`
		Sections []*DocumentSection `json:"sections"`
		// Summary  string             `json:"summary"`
		ID string `json:"_id,omitempty"`
	} `json:"latestDocument"`
	ID             string    `json:"_id"`
	Owners         []string  `json:"owners"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	Rvn            int       `json:"rvn"`
	V              int       `json:"__v"`
	DraftDocuments []struct {
		Document struct {
			// Summary  string             `json:"summary"`
			Updated  time.Time          `json:"updated"`
			Sections []*DocumentSection `json:"sections"`
			ID       string             `json:"_id,omitempty"`
		} `json:"document"`
		VersionTimestamp time.Time `json:"versionTimestamp"`
		ID               string    `json:"_id"`
	} `json:"draftDocuments"`
}

type Document struct {
	Data struct {
		Doc        *DocumentBody `json:"doc,omitempty"`
		PreviewDoc *DocumentBody `json:"previewDoc,omitempty"`
	} `json:"data"`
}

func (d *Document) ID() string {
	if d.Data.Doc == nil {
		return d.Data.PreviewDoc.ID
	} else {
		return d.Data.Doc.ID
	}
}

type DocumentSection struct {
	DatabaseExercise *DatabaseExercise `json:"databaseExercise"`
	CodeExercise     *CodeExercise     `json:"codeExercise,omitempty"`
	Markdown         *Markdown         `json:"markdown,omitempty"`
	Quiz             *Quiz             `json:"quiz,omitempty"`
	Slider           *Slider           `json:"slider"`
	ProductList      []any             `json:"productList"`
	SectionID        string            `json:"sectionId"`
	SectionType      string            `json:"sectionType"`
	ID               string            `json:"_id"`
	Collapse         []any             `json:"collapse"`
	Image            *Image            `json:"image,omitempty"`
	Video            *Video            `json:"video,omitempty"`
}

type DatabaseExercise struct {
	MakeReadOnly           bool               `json:"makeReadOnly"`
	ShowExecuteButton      bool               `json:"showExecuteButton"`
	ShowSubmitButton       bool               `json:"showSubmitButton"`
	ShowRunButton          bool               `json:"showRunButton"`
	ShowSolutionButton     bool               `json:"showSolutionButton"`
	ShowSolutionCode       bool               `json:"showSolutionCode"`
	SkipOutputVerification bool               `json:"skipOutputVerification"`
	Script                 []*DatabaseScript  `json:"script"`
	DatabaseSchema         []*DatabaseSchema  `json:"databaseSchema"`
	ExampleTestData        []*DatabaseExample `json:"exampleTestData"`
	InternalTestData       []*DatabaseTest    `json:"internalTestData"`
}

type DatabaseScript struct {
	Language        string `json:"language"`
	SolutionQueries string `json:"solutionQueries"`
	UserQueries     string `json:"userQueries"`
	ID              string `json:"_id"`
}

type DatabaseSchema struct {
	Type   string           `json:"type"`
	Name   string           `json:"name"`
	Fields []*DatabaseField `json:"fields"`
	ID     string           `json:"_id"`
}

type DatabaseField struct {
	Name string `json:"name"`
	Type string `json:"type"`
	ID   string `json:"_id"`
}

type DatabaseExample struct {
	ExpectedOutput struct {
		Name    string     `json:"name"`
		Headers []string   `json:"headers"`
		Rows    [][]string `json:"rows"`
	} `json:"expectedOutput"`
	Input []struct {
		Name    string     `json:"name"`
		Headers []string   `json:"headers"`
		Rows    [][]string `json:"rows"`
		ID      string     `json:"_id"`
	} `json:"input"`
	ID string `json:"_id"`
}

type DatabaseTest struct {
	ExpectedOutput struct {
		Name    string   `json:"name"`
		Headers []string `json:"headers"`
		Rows    [][]int  `json:"rows"`
	} `json:"expectedOutput"`
	Input []struct {
		Name    string   `json:"name"`
		Headers []string `json:"headers"`
		Rows    [][]int  `json:"rows"`
		ID      string   `json:"_id"`
	} `json:"input"`
	ID string `json:"_id"`
}

type CodeExercise struct {
	Code                     []*Code         `json:"code"`
	ExampleTestCases         []*CodeTestCase `json:"exampleTestCases"`
	InternalTestCases        []*CodeTestCase `json:"internalTestCases"`
	MakeReadOnly             bool            `json:"makeReadOnly"`
	ShowExecuteButton        bool            `json:"showExecuteButton"`
	ShowSubmitButton         bool            `json:"showSubmitButton"`
	ShowRunButton            bool            `json:"showRunButton"`
	ShowSolutionButton       bool            `json:"showSolutionButton"`
	ExecuteWithCustomClasses bool            `json:"executeWithCustomClasses"`
	ShowSolutionCode         bool            `json:"showSolutionCode"`
	SkipOutputVerification   bool            `json:"skipOutputVerification"`
}

type CodeTestCase struct {
	Input          string `json:"input"`
	ExpectedOutput string `json:"expectedOutput"`
	ID             string `json:"_id"`
}

type Code struct {
	Language              string `json:"language"`
	UserCodeHighlight     []any  `json:"userCodeHighlight"`
	SolutionCode          string `json:"solutionCode"`
	SolutionCodeHighlight []any  `json:"solutionCodeHighlight"`
	ID                    string `json:"_id"`
	UserCode              string `json:"userCode,omitempty"`
	DriverCode            string `json:"driverCode,omitempty"`
}

type Quiz struct {
	// Title                     string        `json:"title"`
	Statement                 string        `json:"statement"`
	Options                   []*QuizOption `json:"options"`
	AreMultipleCorrectAnswers bool          `json:"areMultipleCorrectAnswers"`
	IsHiddenAnswerQuiz        bool          `json:"isHiddenAnswerQuiz"`
}

type QuizOption struct {
	Text        string `json:"text"`
	Explanation string `json:"explanation"`
	IsCorrect   bool   `json:"isCorrect"`
	ID          string `json:"_id"`
}

type Markdown struct {
	Text string `json:"text"`
}

type Video struct {
	MediaLink string `json:"mediaLink"`
	Caption   string `json:"caption"`
}

type Slider struct {
	Title  string   `json:"title,omitempty"`
	Images []*Image `json:"images"`
}

type Image struct {
	File    *File  `json:"file"`
	Caption string `json:"caption"`
	Scale   int    `json:"scale"`
}

type File struct {
	Store     string        `json:"store"`
	ID        string        `json:"id"`
	MediaLink string        `json:"mediaLink"`
	Size      int           `json:"size"`
	Md5Hash   string        `json:"md5Hash"`
	Metadata  *FileMetadata `json:"metadata"`
}

type FileMetadata struct {
	Kind                    string    `json:"kind"`
	ID                      string    `json:"id"`
	SelfLink                string    `json:"selfLink"`
	MediaLink               string    `json:"mediaLink"`
	Name                    string    `json:"name"`
	Bucket                  string    `json:"bucket"`
	Generation              string    `json:"generation"`
	Metageneration          string    `json:"metageneration"`
	ContentType             string    `json:"contentType"`
	StorageClass            string    `json:"storageClass"`
	Size                    int       `json:"size"`
	Md5Hash                 string    `json:"md5Hash"`
	Crc32C                  string    `json:"crc32c"`
	Etag                    string    `json:"etag"`
	TimeCreated             time.Time `json:"timeCreated"`
	Updated                 time.Time `json:"updated"`
	TimeStorageClassUpdated time.Time `json:"timeStorageClassUpdated"`
	Metadata                struct {
		DocumentID       string `json:"documentId"`
		OriginalFileName string `json:"originalFileName"`
	} `json:"metadata"`
}

type Courses struct {
	PageProps struct {
		SingleCourse []*Course `json:"singleCourse"`
		PathCourse   []*Course `json:"pathCourse"`
		Courses      []struct {
			InternalID string `json:"_id"`
			Title      string `json:"title"`
			ID         string `json:"id"`
		} `json:"courses"`
		SubscriptionProduct []*Course `json:"subscriptionProduct"`
		BootcampProducts    []*Course `json:"bootcampProducts"`
	} `json:"pageProps"`
	NSSG bool `json:"__N_SSG"`
}

func (c *Courses) CourseByTitle(title string) (*Course, bool) {
	for _, course := range c.PageProps.SingleCourse {
		if course.Title == title {
			return course, true
		}
	}

	return nil, false
}

type Lessons struct {
	Data []*Lesson `json:"data"`
}
