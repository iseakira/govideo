import logging

from fastapi import FastAPI

from app.controllers import health, video
from settings import settings

logging.basicConfig(filename=settings.log_file, level=logging.INFO)

app = FastAPI(
    title="MS Recommendation API",
)

app.include_router(health.router)
app.include_router(video.router)

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=settings.port)

