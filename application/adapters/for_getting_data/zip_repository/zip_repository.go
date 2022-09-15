package zip_repository

import (
	"allureTool/application/adapters/for_getting_data/csv_repository"
	"allureTool/application/domain"
	"log"
)

type ZipRepository struct {
	Archive    ZipArchive
	TargetFile string
}

func (r ZipRepository) Retrieve(project string) domain.ExecutionData {
	err := r.Archive.ExtractFile(r.TargetFile)
	if err != nil {
		log.Fatalf("%#v", err)
	}

	csv := csv_repository.MakeCSVRepositoryForProjectAndFile(r.Archive.Fs, project, "/tmp/"+r.TargetFile)

	return csv.Retrieve(project)
}

func MakeZipRepositoryFromArchive(archive ZipArchive) ZipRepository {
	return ZipRepository{
		Archive:    archive,
		TargetFile: "allure-report/data/behaviors.csv",
	}
}
