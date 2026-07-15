import datetime
from sqlalchemy import Column, DateTime, Float, Integer, String, ForeignKey
from sqlalchemy.schema import UniqueConstraint

from app.models.db import Base


class Rate(Base):
  __tablename__ = "rates"
  __table_args__=(UniqueConstraint("user_id","video_id"),)

  id = Column(Integer, primary_key=True, nullable=False)
  user_id = Column(String(255), nullable=True)
  video_id = Column(ForeignKey("videos.id", ondelete="CASCADE"), nullable=False)
  value = Column(Float, nullable=True, server_default="3.0")
  created_at = Column(DateTime, default=datetime.datetime.utcnow)
  updated_at = Column(DateTime, default=datetime.datetime.utcnow)


  def __init__(self, user_id, video_id, value):
        self.user_id = user_id
        self.video_id = video_id
        self.value = value

