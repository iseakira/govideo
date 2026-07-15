import logging
from fastapi import APIRouter


router = APIRouter(prefix="/api/v1", tags=["Health"])
logger = logging.getLogger(__name__)


@router.get(
    "/health",
    response_model=dict,
    summary="Health check",
)
def health():
    return {"message": "ok"}