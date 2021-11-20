package repository

import (
	"context"

	"github.com/grupogit/RMNGR002001/Backend/domain/model"
)

type ProjectRepository interface {
	GetProjects(ctx context.Context) ([]*model.ProjectCard, error)
	GetProjectByID(ctx context.Context, id int) (*model.ProjectCardDetail, error)
	CreateProject(ctx context.Context, object *model.Project) (bool, error)
	GetResources(ctx context.Context) ([]*model.UserProjectDetail, error)
	GetLeaders(ctx context.Context) ([]*model.UserProjectDetail, error)
	AllocateResource(ctx context.Context, member *model.Member) (bool, error)
	AllocateLeader(ctx context.Context, member *model.Member) (bool, error)
	DeleteResource(ctx context.Context, id int) (bool, error)
	FinishProject(ctx context.Context, id int) error
	UpdateProject(ctx context.Context, object *model.Project) (bool, error) // Pendiente
}
