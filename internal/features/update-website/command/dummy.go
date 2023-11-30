package command

import (
	"context"

	updatewebsite "github.com/GymSquad/archive-api/internal/features/update-website"
)

type dummyCommand struct{}

var _ updatewebsite.UpdateWebisteCommand = (*dummyCommand)(nil)

func NewDummyCommand() *dummyCommand {
	return &dummyCommand{}
}

// Execute implements updatewebsite.UpdateWebisteCommand.
func (*dummyCommand) Execute(context.Context, updatewebsite.UpdateWebsitePayload) (updatewebsite.UpdatedWebsite, error) {
	return updatewebsite.UpdatedWebsite{
		ID:   "clpl86hkn0000356qz6mb44lv",
		Name: "交通大學圖書館",
		Url:  "https://new.lib.nctu.edu.tw/",
		Affiliations: []updatewebsite.UpdatedAffiliation{
			{
				CampusID:       "clpl85gua0000356ribtcvv9d",
				CampusName:     "交大相關",
				DepartmentID:   "clpl862ag0000356qdt2gy68z",
				DepartmentName: "行政單位",
				OfficeID:       "clpl869zn0000356qjk6mc5zu",
				OfficeName:     "圖書館",
			},
		},
	}, nil
}
