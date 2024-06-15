from fastapi import FastAPI

from core.src.router.schema.schemas import Playlist, SongResponse
from core.src.router.scripts.predictor import Predictor

from versioning.versioning import Versioner


app = FastAPI(
	title="ML Core",
	description="Machine Learning Core API",
	version="0.1.0",
)


@app.on_event("startup")
async def startup_event():
	with open("core/src/router/openapi.json", "w") as file:
		import json
		from fastapi.openapi.utils import get_openapi
		openapi_schema = get_openapi(
			title="ML Core",
			version="0.1.0",
			routes=app.routes
		)
		json.dump(openapi_schema, file)


@app.post("/predict", response_model=SongResponse)
async def predict(playlist: Playlist) -> SongResponse:
	return SongResponse(id=str(Predictor().predict(playlist)["id"]))


@app.post("/commit_model")
async def commit_new_model():
	Versioner().commit()
	return {"message": "Model committed successfully!"}
