import logging
from datetime import datetime
from typing import List, Optional

from fastapi import APIRouter, HTTPException, Query
from pydantic import BaseModel, Field

from app.models.video import Video

router = APIRouter(prefix="/api/v1", tags=["Videos"])
logger = logging.getLogger(__name__)



class VideoResponse(BaseModel):
    id: int
    userId: str = Field(alias="user_id")
    title: str
    createdAt: Optional[datetime] = Field(None, alias="created_at")
    updatedAt: Optional[datetime] = Field(None, alias="updated_at")

    model_config = {"from_attributes": True, "populate_by_name": True}


@router.get(
    "/videos/recommended",
    response_model=List[VideoResponse],
    summary="Returns recommended videos",
)
def get_recommended_videos(
    user_id: str = Query(..., description="User ID"),
    limit: int = Query(1000, ge=1, description="Max number of results"),
):
    videos = Video.find_all_recommended(user_id=user_id, limit=limit)
    return videos