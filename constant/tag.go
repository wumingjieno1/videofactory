package constant

type Tag uint64

const (
	TagVideoEmpty              Tag = 0  // 视频空
	TagVideoTitle              Tag = 1  // 视频标题
	TagArticleTitle            Tag = 2  // 文章原标题
	TagArticleTranslateTitle   Tag = 3  // 文章翻译标题
	TagArticleContent          Tag = 4  // 文章原内容
	TagArticleTranslateContent Tag = 5  // 文章翻译内容
)

