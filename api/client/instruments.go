package client

import (
	"context"
	"database/sql"
	"github.com/jannyjacky1/barmen/protogen"
	"github.com/jannyjacky1/barmen/tools"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
)

type InstrumentsServer struct {
	App tools.App
}

func (s *InstrumentsServer) GetInstrumentById(ctx context.Context, request *protogen.InstrumentRequest) (*protogen.InstrumentResponse, error) {

	var instrument protogen.InstrumentResponse

	err := s.App.Db.GetContext(ctx, &instrument, "SELECT tbl_instruments.id, name, coalesce(name_en, '') AS nameEn, COUNT(DISTINCT tbl_cocktails_to_tbl_instruments.cocktail_id) AS info, coalesce(description, '') AS description, coalesce(tbl_files.filepath, '') AS icon FROM tbl_instruments LEFT JOIN tbl_files ON tbl_files.id = tbl_instruments.img_id LEFT JOIN tbl_cocktails_to_tbl_instruments ON tbl_cocktails_to_tbl_instruments.instrument_id = tbl_instruments.id WHERE tbl_instruments.id = $1 GROUP BY tbl_instruments.id, tbl_files.filepath", request.Id)

	if err != nil {
		code := codes.Internal
		if err == sql.ErrNoRows {
			code = codes.NotFound
		}
		return &instrument, status.Error(code, err.Error())
	}

	cnt, _ := strconv.Atoi(instrument.Info)
	instrument.Info = "Используется для " + instrument.Info + " " + tools.GetWord(cnt, "коктейля", "коктейлей", "коктейлей")

	return &instrument, nil
}
