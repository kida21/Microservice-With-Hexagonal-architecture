package adapter

import (
	"context"

	"github.com/kida21/authservice/internal/domain"
)
type Adapter struct{}
func NewAdapter()(*Adapter){
	return &Adapter{}
}
func (a *Adapter) Login(ctx context.Context,input *domain.Credential)(string,error){
	return "",nil
}