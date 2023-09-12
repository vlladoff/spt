package file_storage

import (
	"bufio"
	"encoding/json"
	"errors"
	"github.com/gocarina/gocsv"
	"github.com/vlladoff/spt/internal/predict"
	"io"
	"os"
	"path"
)

type FileStorage struct{}

func (fl FileStorage) GetLtvData(sourcePath string) (*[]*predict.DataToPredict, error) {
	data := new([]*predict.DataToPredict)
	ext := path.Ext(sourcePath)

	file, err := os.OpenFile(sourcePath, os.O_RDWR, os.ModePerm)
	if err != nil {
		return data, err
	}
	defer file.Close()

	type TempDataToPredict struct {
		CampaignId                               string `csv:"CampaignId" json:"CampaignId"`
		Country                                  string `csv:"Country" json:"Country"`
		Users                                    int32  `json:"Users"`
		Ltv1, Ltv2, Ltv3, Ltv4, Ltv5, Ltv6, Ltv7 float64
	}
	tempData := new([]*TempDataToPredict)

	switch ext {
	case predict.FileCsv:
		err = gocsv.UnmarshalFile(file, tempData)
	case predict.FileJson:
		stat, err := file.Stat()
		if err != nil {
			return data, err
		}

		bs := make([]byte, stat.Size())
		_, err = bufio.NewReader(file).Read(bs)
		if err != nil && err != io.EOF {
			return data, err
		}

		err = json.Unmarshal(bs, tempData)
	default:
		err = errors.New("wrong file type")
	}

	for _, val := range *tempData {
		var ltv1to7Arr [7]float64
		ltv1to7Arr[0] = val.Ltv1
		ltv1to7Arr[1] = val.Ltv2
		ltv1to7Arr[2] = val.Ltv3
		ltv1to7Arr[3] = val.Ltv4
		ltv1to7Arr[4] = val.Ltv5
		ltv1to7Arr[5] = val.Ltv6
		ltv1to7Arr[6] = val.Ltv7

		dataVal := predict.DataToPredict{
			CampaignId: val.CampaignId,
			Country:    val.Country,
			Users:      val.Users,
			Ltv1to7:    ltv1to7Arr,
		}

		*data = append(*data, &dataVal)
	}

	return data, err
}
