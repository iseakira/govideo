import datetime

import pandas as pd
from sqlalchemy import Column, DateTime, Integer, String
from surprise import SVD, Dataset, NormalPredictor, Reader
from surprise.model_selection import cross_validate

from app.models.rate import Rate
from app.models.db import Base, database


class Video(Base):
  __tablename__ = "videos"
  id = Column(Integer, primary_key=True, nullable=False)
  user_id = Column(String(255), nullable=True)
  title = Column(String(255), nullable=False)
  created_at = Column(DateTime, nullable=False, default=datetime.datetime.utcnow)
  updated_at = Column(DateTime, nullable=False, default=datetime.datetime.utcnow)

  def __init__(self, user_id, title, id=None, created_at=None, updated_at=None):
        self.user_id = user_id
        self.title = title
        if id:
            self.id = id
        if created_at:
            self.created_at = created_at
        if updated_at:
            self.updated_at = updated_at


  @staticmethod
  def predict_recommend(
            df: pd.DataFrame,
            dataset_columns: list,
            item_id: int,
            user_id: str,
            rating_scale_st: int = 1,
            rating_scale_end: int = 5,
            cross_validate_value: int = 2,
    ):
      reader = Reader(rating_sclae=(rating_scale_st,rating_scale_end))
      data = Dataset.load_from_df(df[dataset_columns],reader)

      try:
           cross_validate(NormalPredictor(), data, cv=cross_validate_value)
      except ValueError:
          return None

      svd = SVD()

      trainset = data.build_full_trainset()
      svd.fit(trainset)

      predict_df = df.copy()
      predict_df["Predicted_Score"] = predict_df[item_id].apply(
            lambda x: svd.predict(user_id, x).est
        )
      predict_df = predict_df.sort_values(by=["Predicted_Score"], ascending=False)
      predict_df = predict_df.drop_duplicates(subset=item_id)
      return predict_df


  @staticmethod
  def find_all_recommned(user_id:str,limit: int=1000):
      import logging
      session = database.connect_db

      try:
          query = session.query(Video, Rate).filter(Video.id == Rate.video_id).limit(limit).statement
          df = pd.read_sql(query, session.bind)

          if df.empty:
                return []

          df.columns = [
                "id",
                "user_id",
                "title",
                "created_at",
                "updated_at",
                "rate_id",
                "rate_user_id",
                "rate_video_id",
                "rate_value",
                "rate_created_at",
                "rate_updated_at",
            ]

          recommended_df = Video.predict_recommend(
                df=df,
                dataset_columns=["rate_user_id", "rate_video_id", "rate_value"],
                item_id="rate_video_id",
                user_id=user_id,
            )

          if recommended_df is None:
                return []

          recommended_df = recommended_df.head(limit)

          recommended_videos = []
          for _, row in recommended_df.iterrows():
                video = Video(
                    id=row["id"],
                    user_id=row["user_id"],
                    title=row["title"],
                    created_at=row["created_at"],
                    updated_at=row["updated_at"]
                )
                recommended_videos.append(video)
          return recommended_videos
      except Exception as e:
            logging.error("find_all_recommended failed: %s", e)
            return []
      finally:
            session.close()


  @property
  def json(self):
        return {
            "id": self.id,
            "user_id": self.user_id,
            "title": self.title,
            "created_at": self.created_at,
            "updated_at": self.updated_at,
        }















