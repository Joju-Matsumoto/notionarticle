package notionarticle

import (
	"testing"

	"github.com/jomei/notionapi"
	"github.com/stretchr/testify/require"
)

func TestArticleJson(t *testing.T) {
	want := Article{
		Page: notionapi.Page{
			ID: "hello",
		},
		Blocks: []notionapi.Block{
			notionapi.ParagraphBlock{
				BasicBlock: notionapi.BasicBlock{
					Type:        notionapi.BlockTypeParagraph,
					HasChildren: false,
				},
				Paragraph: notionapi.Paragraph{
					RichText: []notionapi.RichText{
						{
							Type: notionapi.ObjectTypeText,
							Text: &notionapi.Text{
								Content: "hello",
							},
							PlainText: "hello",
						},
					},
				},
			},
		},
	}

	err := saveJson("./testdata/sample_article.json", &want)
	require.NoError(t, err)

	var got Article
	err = loadJson("./testdata/sample_article.json", &got)
	require.NoError(t, err)
}
