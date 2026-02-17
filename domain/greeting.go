package domain

import (
	"time"
    "errors"
    "regexp"
    "unicode/utf8"
)

type GreetingInput struct {
	Name string 
	Time time.Time 
}

type GreetingOutput struct {
    Message string
}

// 挨拶を判定するロジック（ドメインモデルのメソッド）
func (i *GreetingInput) GenerateMessage() string {
    jst := time.FixedZone("Asia/Tokyo", 9*60*60)
    localTime := i.Time.In(jst)
    hour := localTime.Hour()
    var greeting string
    switch {
    case hour >= 5 && hour < 12:
        greeting = "おはようございます"
    case hour >= 12 && hour < 18:
        greeting = "こんにちは"
    default:
        greeting = "こんばんは"
    }
    return greeting + "、" + i.Name + "！"
}

func (i *GreetingInput) Validate() error {
    //未入力チェック
    if i.Name == "" {
        return errors.New("名前が入力されていません")
    }

    //文字数チェック (20文字以内)
    if utf8.RuneCountInString(i.Name) > 20 {
        return errors.New("名前は20文字以内で入力してください")
    }

    //全角文字チェック
    isFullWidth := regexp.MustCompile(`^[^\x01-\x7E\xA1-\xDF]+$`).MatchString(i.Name)
    if !isFullWidth {
        return errors.New("名前は全角文字で入力してください")
    }

    return nil
}