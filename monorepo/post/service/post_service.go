package service

import (
	"fmt"
	"v1/post/model"
	"v1/post/repository"
)

type PostService interface {
	GetFeed(accontId uint64) ([]model.Post, error)
	GetPosts(accountId uint64, communityId uint64) ([]model.Post, error)
	GetPostByName() []model.Post
	CreatePost(userId uint64, postBody model.PostDTO) error
	UpdatePost(postID uint64, userID uint64, postBody model.PostDTO) error
	DeletePost(postID uint64, userID uint64) error
}

type postService struct {
	postRepository repository.PostRepository
}

func NewPostService(postRepository repository.PostRepository) PostService {
	return &postService{postRepository: postRepository}
}

func (p *postService) GetPosts(accountId uint64, communityId uint64) ([]model.Post, error) {
	if accountId != 0 {
		posts, err := p.postRepository.FindUserPosts(accountId)
		if err != nil {
			return nil, err
		}

		return posts, nil
	}

	posts, err := p.postRepository.FindCommunityPosts(communityId)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (p *postService) CreatePost(userId uint64, postBody model.PostDTO) error {
	if err := postBody.Prepare(); err != nil {
		return fmt.Errorf("error creating post: %w", err)
	}

	if err := p.postRepository.Create(userId, postBody); err != nil {
		return fmt.Errorf("error creating post: %w", err)
	}

	return nil
}

func (p *postService) GetFeed(accountId uint64) ([]model.Post, error) {
	// if accountId != 0 {
	// 	posts, err := p.postRepository.FindUserPosts(accountId)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	return posts, nil
	// }

	// posts, err := p.postRepository.FindCommunityPosts(communityId)
	// if err != nil {
	// 	return nil, err
	// }

	// return posts, nil
	return nil, nil
}

func (p *postService) GetPostByName() []model.Post {
	return []model.Post{}
}

func (p *postService) UpdatePost(postID uint64, userID uint64, postBody model.PostDTO) error {
	if err := postBody.Prepare(); err != nil {
		return fmt.Errorf("invalid post: %w", err)
	}
	return p.postRepository.Update(postID, userID, postBody)
}

func (p *postService) DeletePost(postID uint64, userID uint64) error {
	return p.postRepository.Delete(postID, userID)
}
