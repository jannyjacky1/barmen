package client

import (
	"context"
	"github.com/jannyjacky1/barmen/proto"
	"log"
)

type DictionariesServer struct {
}

func (s *DictionariesServer) GetDictionaries(ctx context.Context, request *proto.DictionariesRequest) (*proto.DictionariesResponse, error) {

	complicationLevels, err := getComplicationLevels()
	if err != nil {
		log.Fatalln(err)
	}

	fortressLevels, err := getFortressLevels()
	if err != nil {
		log.Fatalln(err)
	}

	volumes, err := getVolumes()
	if err != nil {
		log.Fatalln(err)
	}

	response := proto.DictionariesResponse{ComplicationLevels: complicationLevels, FortressLevels: fortressLevels, Volumes: volumes}
	return &response, nil
}

func getComplicationLevels() ([]*proto.Dictionary, error) {
	items := make([]*proto.Dictionary, 3)
	items[0] = &proto.Dictionary{Id: 1, Name: "Легко"}
	items[1] = &proto.Dictionary{Id: 2, Name: "Средне"}
	items[2] = &proto.Dictionary{Id: 3, Name: "Сложно"}

	return items, nil
}

func getFortressLevels() ([]*proto.Dictionary, error) {
	items := make([]*proto.Dictionary, 4)
	items[0] = &proto.Dictionary{Id: 1, Name: "Безалкогольный"}
	items[1] = &proto.Dictionary{Id: 2, Name: "Легкий"}
	items[2] = &proto.Dictionary{Id: 3, Name: "Средний"}
	items[3] = &proto.Dictionary{Id: 4, Name: "Сложный"}

	return items, nil
}

func getVolumes() ([]*proto.Dictionary, error) {
	items := make([]*proto.Dictionary, 4)
	items[0] = &proto.Dictionary{Id: 1, Name: "до 60 мл"}
	items[1] = &proto.Dictionary{Id: 2, Name: "60-120 мл"}
	items[2] = &proto.Dictionary{Id: 3, Name: "120-250 мл"}
	items[3] = &proto.Dictionary{Id: 4, Name: "более 250 мл"}

	return items, nil
}
