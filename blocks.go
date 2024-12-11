package notionarticle

import "github.com/jomei/notionapi"

// utility functions for notionapi.Block, notionapi.Blocks

func Walk(blocks notionapi.Blocks, f func(b notionapi.Block) error) error {
	for _, b := range blocks {
		if err := f(b); err != nil {
			return err
		}
		if b.GetHasChildren() {
			if cPtr := Children(b); cPtr != nil {
				if err := Walk(*cPtr, f); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func Children(b notionapi.Block) *notionapi.Blocks {
	switch b := b.(type) {
	case *notionapi.BulletedListItemBlock:
		return &b.BulletedListItem.Children
	case *notionapi.CalloutBlock:
		return &b.Callout.Children
	// case *notionapi.ChildDatabaseBlock:
	// case *notionapi.ChildPageBlock:
	case *notionapi.ColumnBlock:
		return &b.Column.Children
	case *notionapi.ColumnListBlock:
		return &b.ColumnList.Children
	case *notionapi.Heading1Block:
		return &b.Heading1.Children
	case *notionapi.Heading2Block:
		return &b.Heading2.Children
	case *notionapi.Heading3Block:
		return &b.Heading3.Children
	case *notionapi.NumberedListItemBlock:
		return &b.NumberedListItem.Children
	case *notionapi.ParagraphBlock:
		return &b.Paragraph.Children
	case *notionapi.QuoteBlock:
		return &b.Quote.Children
	case *notionapi.SyncedBlock:
		return &b.SyncedBlock.Children
	case *notionapi.TableBlock:
		return &b.Table.Children
	case *notionapi.TemplateBlock:
		return &b.Template.Children
	case *notionapi.ToDoBlock:
		return &b.ToDo.Children
	case *notionapi.ToggleBlock:
		return &b.Toggle.Children
	}
	return nil
}
