package usecase

import (
	"context"

	"github.com/grupogit/RMNGR002001/Backend/domain/repository"

	"github.com/grupogit/RMNGR002001/Backend/domain/model"
)

type ProjectUseCase interface {
	GetProjects(ctx context.Context) ([]*model.ProjectCard, error)
	GetProjectByID(ctx context.Context, id int) (*model.ProjectCardDetail, error)
	CreateProject(ctx context.Context, object *model.Project) (bool, error)
	GetResources(ctx context.Context) ([]*model.UserProjectDetail, error)
	GetLeaders(ctx context.Context) ([]*model.UserProjectDetail, error)
	AllocateResource(ctx context.Context, member *model.Member) (bool, error)
	AllocateLeader(ctx context.Context, member *model.Member) (bool, error)
	DeleteResource(ctx context.Context, id int) (bool, error)
	FinishProject(ctx context.Context, id int) error
	UpdateProject(ctx context.Context, object *model.Project) (bool, error)
}

type projectUseCase struct {
	repository.ProjectRepository
}

func NewProjectUseCase(r repository.ProjectRepository) ProjectUseCase {
	return &projectUseCase{r}
}

func (u *projectUseCase) GetProjects(ctx context.Context) ([]*model.ProjectCard, error) {
	return u.ProjectRepository.GetProjects(ctx)
}

func (u *projectUseCase) GetProjectByID(ctx context.Context, id int) (*model.ProjectCardDetail, error) {
	return u.ProjectRepository.GetProjectByID(ctx, id)
}

func (u *projectUseCase) CreateProject(ctx context.Context, object *model.Project) (bool, error) {
	return u.ProjectRepository.CreateProject(ctx, object)
}

func (u *projectUseCase) GetResources(ctx context.Context) ([]*model.UserProjectDetail, error) {
	return u.ProjectRepository.GetResources(ctx)
}

func (u *projectUseCase) GetLeaders(ctx context.Context) ([]*model.UserProjectDetail, error) {
	return u.ProjectRepository.GetLeaders(ctx)
}

func (u *projectUseCase) AllocateResource(ctx context.Context, member *model.Member) (bool, error) {
	return u.ProjectRepository.AllocateResource(ctx, member)
}

func (u *projectUseCase) AllocateLeader(ctx context.Context, member *model.Member) (bool, error) {
	return u.ProjectRepository.AllocateLeader(ctx, member)
}

func (u *projectUseCase) DeleteResource(ctx context.Context, id int) (bool, error) {
	return u.ProjectRepository.DeleteResource(ctx, id)
}

func (u *projectUseCase) FinishProject(ctx context.Context, id int) error {
	return u.ProjectRepository.FinishProject(ctx, id)
}

func (u *projectUseCase) UpdateProject(ctx context.Context, object *model.Project) (bool, error) {
	return u.ProjectRepository.UpdateProject(ctx, object)
}
