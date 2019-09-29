package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/AlexeySemin/test/golang-service/repositories"
	"github.com/AlexeySemin/test/golang-service/response"
	"github.com/jinzhu/gorm"
)

type SSAController struct {
	db         *gorm.DB
	repository *repositories.SSARepository
}

type MapYearMonthIndex struct {
	Year  int
	Month int
}

type DataByMonth struct {
	Date      string  `json:"date"`
	AvgRating float64 `json:"avgRating"`
	MinRating int     `json:"minRating"`
	MaxRating int     `json:"maxRating"`
	CountNews int     `json:"countNews"`
}

type RatingByMonth struct {
	Count    int
	Sum      int
	Min      int
	Max      int
	isMinSet bool
}

func NewSSAController(db *gorm.DB) *SSAController {
	repository := repositories.NewSSARepository(db)
	return &SSAController{db, repository}
}

func (ssac *SSAController) GetMinMaxAvgRating(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	useRows := r.FormValue("use_rows")
	var minMaxAvgResp *response.MinMaxAvgRating
	var err error

	if useRows == "" || useRows == "false" {
		minMaxAvgResp, err = ssac.getMinMaxAvgRating()
	} else {
		minMaxAvgResp, err = ssac.getMinMaxAvgRatingUsingRows()
	}

	if err != nil {
		response.Send(w, nil, err.Error(), http.StatusInternalServerError)
		return
	}

	end := time.Now()
	logResp := response.NewLog(start, end)
	resp := struct {
		response.MinMaxAvgRating
		response.Log
	}{*minMaxAvgResp, *logResp}

	response.Send(w, resp, "", http.StatusOK)
}

func (ssac *SSAController) getMinMaxAvgRating() (*response.MinMaxAvgRating, error) {
	news, err := ssac.repository.GetNews()
	if err != nil {
		return nil, err
	}

	min := 0
	max := 0
	sumRating := 0
	count := 0
	avg := 0.0

	for _, oneNews := range news {
		if oneNews.Rating < min {
			min = oneNews.Rating
		}
		if oneNews.Rating > max {
			max = oneNews.Rating
		}
		sumRating += oneNews.Rating
		count++
	}

	if count > 0 {
		avg = float64(sumRating) / float64(count)
	}

	return &response.MinMaxAvgRating{
		Min: min,
		Max: max,
		Avg: avg,
	}, nil
}

func (ssac *SSAController) getMinMaxAvgRatingUsingRows() (*response.MinMaxAvgRating, error) {
	rows, err := ssac.repository.GetRatingsRows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	min := 0
	max := 0
	sumRating := 0
	count := 0
	avg := 0.0
	var rating int

	for rows.Next() {
		if err := rows.Scan(&rating); err != nil {
			log.Print(err)
		}
		if rating < min {
			min = rating
		}
		if rating > max {
			max = rating
		}
		sumRating += rating
		count++
	}

	if count > 0 {
		avg = float64(sumRating) / float64(count)
	}

	return &response.MinMaxAvgRating{
		Min: min,
		Max: max,
		Avg: avg,
	}, nil
}

func (ssac *SSAController) GetPerMonthJSONData(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	useRows := r.FormValue("use_rows")
	var perMonthJSONResp *response.PerMonthJSONData
	var err error

	if useRows == "" || useRows == "false" {
		perMonthJSONResp, err = ssac.getPerMonthJSONData()
	} else {
		perMonthJSONResp, err = ssac.getPerMonthJSONDataUsingRows()
	}

	if err != nil {
		response.Send(w, nil, err.Error(), http.StatusInternalServerError)
		return
	}

	end := time.Now()
	logResp := response.NewLog(start, end)
	resp := struct {
		response.PerMonthJSONData
		response.Log
	}{*perMonthJSONResp, *logResp}

	response.Send(w, resp, "", http.StatusOK)
}

func (ssac *SSAController) getPerMonthJSONData() (*response.PerMonthJSONData, error) {
	news, err := ssac.repository.GetNews()
	if err != nil {
		return nil, err
	}

	ratingByMonth := map[MapYearMonthIndex]RatingByMonth{}
	for _, oneNews := range news {
		mapIndex := MapYearMonthIndex{oneNews.CreatedAt.Year(), int(oneNews.CreatedAt.Month())}
		max := ratingByMonth[mapIndex].Max
		if oneNews.Rating > max {
			max = oneNews.Rating
		}
		min := ratingByMonth[mapIndex].Min
		isMinSet := ratingByMonth[mapIndex].isMinSet
		if oneNews.Rating < min || !isMinSet {
			min = oneNews.Rating
			isMinSet = true
		}
		ratingByMonth[mapIndex] = RatingByMonth{
			Count:    ratingByMonth[mapIndex].Count + 1,
			Sum:      ratingByMonth[mapIndex].Sum + oneNews.Rating,
			Min:      min,
			Max:      max,
			isMinSet: isMinSet,
		}
	}

	dataByMonth := []*DataByMonth{}
	for key, data := range ratingByMonth {
		dateStr := strconv.Itoa(key.Year) + "-" + fmt.Sprintf("%02d", key.Month) + "-01"
		dataByMonth = append(dataByMonth, &DataByMonth{
			Date:      dateStr,
			AvgRating: float64(data.Sum) / float64(data.Count),
			MinRating: data.Min,
			MaxRating: data.Max,
			CountNews: data.Count,
		})
	}
	dataByMonthJSON, err := json.Marshal(&dataByMonth)
	if err != nil {
		return nil, err
	}

	return &response.PerMonthJSONData{
		Data: string(dataByMonthJSON),
	}, nil
}

func (ssac *SSAController) getPerMonthJSONDataUsingRows() (*response.PerMonthJSONData, error) {
	rows, err := ssac.repository.GetRatingsAndDatesRows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rating int
	var createdAt time.Time
	ratingByMonth := map[MapYearMonthIndex]RatingByMonth{}
	for rows.Next() {
		if err := rows.Scan(&rating, &createdAt); err != nil {
			log.Print(err)
		}

		mapIndex := MapYearMonthIndex{createdAt.Year(), int(createdAt.Month())}
		max := ratingByMonth[mapIndex].Max
		if rating > max {
			max = rating
		}
		min := ratingByMonth[mapIndex].Min
		isMinSet := ratingByMonth[mapIndex].isMinSet
		if rating < min || !isMinSet {
			min = rating
			isMinSet = true
		}
		ratingByMonth[mapIndex] = RatingByMonth{
			Count:    ratingByMonth[mapIndex].Count + 1,
			Sum:      ratingByMonth[mapIndex].Sum + rating,
			Min:      min,
			Max:      max,
			isMinSet: isMinSet,
		}
	}

	dataByMonth := []*DataByMonth{}
	for key, data := range ratingByMonth {
		dateStr := strconv.Itoa(key.Year) + "-" + fmt.Sprintf("%02d", key.Month) + "-01"
		dataByMonth = append(dataByMonth, &DataByMonth{
			Date:      dateStr,
			AvgRating: float64(data.Sum) / float64(data.Count),
			MinRating: data.Min,
			MaxRating: data.Max,
			CountNews: data.Count,
		})
	}
	dataByMonthJSON, err := json.Marshal(&dataByMonth)
	if err != nil {
		return nil, err
	}

	return &response.PerMonthJSONData{
		Data: string(dataByMonthJSON),
	}, nil
}
