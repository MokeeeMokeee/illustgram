import { Button } from '@material-ui/core';
import React, { useEffect, useState } from 'react';
import { useMutation } from 'react-apollo';
import { Link } from 'react-router-dom';
import styled from 'styled-components';
import { ChangeDiscloseQuery, DeleteContentQuery } from '../data/mutation';

const Container = styled.div``;

const Content = styled.div`
  display: flex;
  height: 160px;
  width: 100%;
  margin-bottom: 2rem;
  @media (max-width: 768px) {
    flex-direction: column;
    height: 100%;
    border: 1px dotted;
  }
`;

const ImageContainer = styled.div`
  height: 100%;
  width: 100%;
  flex: 0 0 160px;
`;

const Image = styled.img`
  width: 100%;
  height: 100%;
  object-fit: cover;
`;

const InfoContainer = styled.div`
  flex: 1 1 auto;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  padding: 0 1rem;
  @media (max-width: 768px) {
    flex: 0 1 auto;
    padding: 0;
  }
`;

const TitleContainer = styled.div`
  @media (max-width: 768px) {
    display: flex;
    margin: 1rem 0 0.5rem 0;
    height: 30px;
  }
`;

const Title = styled.p`
  text-overflow: ellipsis;
  overflow: hidden;
  white-space: nowrap;
  font-weight: bold;
  font-size: 18px;
  @media (max-width: 768px) {
    font-size: 24px;
    margin-right: 0.5rem;
  }
`;

const Detail = styled.p`
  width: 100%;
  overflow: hidden;
  text-overflow: ellipsis;
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
`;

const InfoContent = styled.div`
  display: flex;
  flex-direction: column;
  @media (max-width: 768px) {
    margin: 0.5rem 0;
  }
`;

const ViewCount = styled.p``;

const LikeCount = styled.p``;

const TagsContainer = styled.div`
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
`;

const TagText = styled.p`
  margin-right: 0.3rem;
  &:hover {
    color: #000000;
    font-weight: bold;
  }
`;

const IsPublicContainer = styled.div`
  flex: 0 0 120px;
  display: flex;
  align-items: center;
  justify-content: center;
  @media (max-width: 768px) {
    flex: 0 1 auto;
    display: none;
  }
`;

const IsPublicContent = styled.div`
  display: flex;
  justify-content: center;
  align-items: center;
`;

const PublicText = styled.p``;

const PublicIcon = styled.svg`
  height: 24px;
  width: 24px;
  margin-right: 0.2rem;
`;

const UploadTimeContainer = styled.div`
  flex: 0 0 100px;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;
  @media (max-width: 768px) {
    display: none;
  }
`;

const UploadTimeText = styled.p``;

const UploadText = styled.p`
  font-size: 10px;
`;

const LinkStyled = styled(Link)`
  text-decoration: none;
  color: inherit;
`;

const FunctionSection = styled.div`
  flex: 0 0 110px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
`;

const RightContent = styled.div`
  display: flex;
  width: 100%;
  @media (max-width: 768px) {
    flex-direction: column;
  }
`;

const IsSpPublicContainer = styled(IsPublicContainer)`
  display: none;
  @media (max-width: 768px) {
    display: block;
  }
`;

const IsUploadTimeContainer = styled(UploadTimeContainer)`
  display: none;
  @media (max-width: 768px) {
    display: block;
    flex: 0 1 auto;
    margin-bottom: 1rem;
  }
`;

/**
 *
 * interface
 */

interface PostContentProps {
  items: {
    content_id: string;
    user_id: string;
    detail: string;
    create_at: string;
    title: string;
    views: number;
    adult: boolean;
    thumbnail_url: string;
    image_index: number;
    like_count: number;
    disclose: boolean;
  };
}

const PostContent: React.FC<PostContentProps> = (props) => {
  const [disclose, isDisclose] = useState<boolean>();
  const [changeDiscloseQuery] = useMutation(ChangeDiscloseQuery);
  const [deleteContentQuery] = useMutation(DeleteContentQuery);

  useEffect(() => {
    isDisclose(props.items.disclose);
  }, [props.items.disclose]);

  const PublicFunc = (contentId: string) => {
    let changeData = { content_id: contentId, disclose: !disclose };
    changeDiscloseQuery({
      variables: { changeData },
    }).then((res) => {
      if (!res.errors) {
        console.log(res);
        isDisclose(!disclose);
        alert(disclose ? '非公開にしました' : '公開しました');
      } else {
        console.log(res.errors);
      }
    });

    console.log(contentId);
    console.log(!disclose);
  };

  const DeleteItem = (contentId: string, title: string) => {
    console.log(contentId);
    let ContentId = { content_id: contentId };
    deleteContentQuery({
      variables: { ContentId },
    }).then((res) => {
      if (!res.errors) {
        console.log(res);
        alert(`${title}を削除しました。`);
        window.location.reload(false);
      } else {
        console.log(res.errors);
      }
    });
  };

  return (
    <Container>
      <Content>
        <ImageContainer>
          <LinkStyled to={`/illustratio/${props.items.content_id}`}>
            <Image src={props.items.thumbnail_url} />
          </LinkStyled>
        </ImageContainer>
        <RightContent>
          <InfoContainer>
            <TitleContainer>
              <Title>{props.items.title}</Title>
              <IsSpPublicContainer>
                {disclose ? (
                  <IsPublicContent>
                    <PublicIcon
                      viewBox='0 0 24 24'
                      fill='none'
                      xmlns='http://www.w3.org/2000/svg'
                    >
                      <path
                        d='M12 4.5C7 4.5 2.73 7.61 1 12C2.73 16.39 7 19.5 12 19.5C17 19.5 21.27 16.39 23 12C21.27 7.61 17 4.5 12 4.5ZM12 17C9.24 17 7 14.76 7 12C7 9.24 9.24 7 12 7C14.76 7 17 9.24 17 12C17 14.76 14.76 17 12 17ZM12 9C10.34 9 9 10.34 9 12C9 13.66 10.34 15 12 15C13.66 15 15 13.66 15 12C15 10.34 13.66 9 12 9Z'
                        fill='black'
                      />
                    </PublicIcon>
                    <PublicText>公開</PublicText>
                  </IsPublicContent>
                ) : (
                  <IsPublicContent>
                    <PublicIcon
                      viewBox='0 0 24 24'
                      fill='none'
                      xmlns='http://www.w3.org/2000/svg'
                    >
                      <path
                        d='M12 7C14.76 7 17 9.24 17 12C17 12.65 16.87 13.26 16.64 13.83L19.56 16.75C21.07 15.49 22.26 13.86 22.99 12C21.26 7.61 16.99 4.5 11.99 4.5C10.59 4.5 9.25 4.75 8.01 5.2L10.17 7.36C10.74 7.13 11.35 7 12 7ZM2 4.27L4.28 6.55L4.74 7.01C3.08 8.3 1.78 10.02 1 12C2.73 16.39 7 19.5 12 19.5C13.55 19.5 15.03 19.2 16.38 18.66L16.8 19.08L19.73 22L21 20.73L3.27 3L2 4.27ZM7.53 9.8L9.08 11.35C9.03 11.56 9 11.78 9 12C9 13.66 10.34 15 12 15C12.22 15 12.44 14.97 12.65 14.92L14.2 16.47C13.53 16.8 12.79 17 12 17C9.24 17 7 14.76 7 12C7 11.21 7.2 10.47 7.53 9.8V9.8ZM11.84 9.02L14.99 12.17L15.01 12.01C15.01 10.35 13.67 9.01 12.01 9.01L11.84 9.02Z'
                        fill='black'
                      />
                    </PublicIcon>
                    <PublicText>非公開</PublicText>
                  </IsPublicContent>
                )}
              </IsSpPublicContainer>
            </TitleContainer>
            <IsUploadTimeContainer>
              <UploadTimeText>
                {props.items.create_at.substr(0, 10).replace(/-/g, '/')}
              </UploadTimeText>
              <UploadText>アップロード</UploadText>
            </IsUploadTimeContainer>
            <Detail>{props.items.detail}</Detail>
            <InfoContent>
              <ViewCount>
                視聴数: {props.items.views.toLocaleString()}
              </ViewCount>
              <LikeCount>
                いいね数: {props.items.like_count.toLocaleString()}
              </LikeCount>
            </InfoContent>
            {/* <TagsContainer>
              {props.items.tags.map((tag, index) => {
                return (
                  <TagText>
                    <LinkStyled to={`/tags/illustratio/${tag}`}>
                      #{tag}
                    </LinkStyled>
                  </TagText>
                );
              })}
            </TagsContainer> */}
          </InfoContainer>
          <UploadTimeContainer>
            <UploadTimeText>
              {props.items.create_at.substr(0, 10).replace(/-/g, '/')}
            </UploadTimeText>
            <UploadText>アップロード</UploadText>
          </UploadTimeContainer>
          <IsPublicContainer>
            {disclose ? (
              <IsPublicContent>
                <PublicIcon
                  viewBox='0 0 24 24'
                  fill='none'
                  xmlns='http://www.w3.org/2000/svg'
                >
                  <path
                    d='M12 4.5C7 4.5 2.73 7.61 1 12C2.73 16.39 7 19.5 12 19.5C17 19.5 21.27 16.39 23 12C21.27 7.61 17 4.5 12 4.5ZM12 17C9.24 17 7 14.76 7 12C7 9.24 9.24 7 12 7C14.76 7 17 9.24 17 12C17 14.76 14.76 17 12 17ZM12 9C10.34 9 9 10.34 9 12C9 13.66 10.34 15 12 15C13.66 15 15 13.66 15 12C15 10.34 13.66 9 12 9Z'
                    fill='black'
                  />
                </PublicIcon>
                <PublicText>公開</PublicText>
              </IsPublicContent>
            ) : (
              <IsPublicContent>
                <PublicIcon
                  viewBox='0 0 24 24'
                  fill='none'
                  xmlns='http://www.w3.org/2000/svg'
                >
                  <path
                    d='M12 7C14.76 7 17 9.24 17 12C17 12.65 16.87 13.26 16.64 13.83L19.56 16.75C21.07 15.49 22.26 13.86 22.99 12C21.26 7.61 16.99 4.5 11.99 4.5C10.59 4.5 9.25 4.75 8.01 5.2L10.17 7.36C10.74 7.13 11.35 7 12 7ZM2 4.27L4.28 6.55L4.74 7.01C3.08 8.3 1.78 10.02 1 12C2.73 16.39 7 19.5 12 19.5C13.55 19.5 15.03 19.2 16.38 18.66L16.8 19.08L19.73 22L21 20.73L3.27 3L2 4.27ZM7.53 9.8L9.08 11.35C9.03 11.56 9 11.78 9 12C9 13.66 10.34 15 12 15C12.22 15 12.44 14.97 12.65 14.92L14.2 16.47C13.53 16.8 12.79 17 12 17C9.24 17 7 14.76 7 12C7 11.21 7.2 10.47 7.53 9.8V9.8ZM11.84 9.02L14.99 12.17L15.01 12.01C15.01 10.35 13.67 9.01 12.01 9.01L11.84 9.02Z'
                    fill='black'
                  />
                </PublicIcon>
                <PublicText>非公開</PublicText>
              </IsPublicContent>
            )}
          </IsPublicContainer>
          <FunctionSection>
            <Button
              color='primary'
              onClick={() => PublicFunc(props.items.content_id)}
            >
              {disclose ? '非公開にする' : '公開する'}
            </Button>

            <Button
              color='secondary'
              onClick={() =>
                DeleteItem(props.items.content_id, props.items.title)
              }
            >
              削除する
            </Button>
          </FunctionSection>
        </RightContent>
      </Content>
    </Container>
  );
};

export default PostContent;
