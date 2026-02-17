package usecase

import (
	"exercise-api-for-pj1/domain"
)

func ExecuteGreeting(input domain.GreetingInput) (domain.GreetingOutput,error) {
    // ドメインロジックを呼び出す
    msg := input.GenerateMessage()
    return domain.GreetingOutput{Message: msg}, nil
}