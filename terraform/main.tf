provider "google" {
  project = "gcp-mig-demo-340020"
  region  = "us-central1"
}

provider "google-beta" {
  project = "gcp-mig-demo-340020"
  region  = "us-central1"
}

resource "google_artifact_registry_repository" "my-repo" {
  provider = google-beta

  location = "us-central1"
  repository_id = "gpc-mig-example"
  description = "example docker repository"
  format = "DOCKER"
}

resource "google_compute_network" "vpc_network" {
  name                    = "vpc-network"
  auto_create_subnetworks = true
}

resource "google_service_account" "service_account" {
  account_id   = "gcp-mig-demo"
  display_name = "gcp mig demo account"
}

resource "google_service_account_iam_binding" "artifact_registry" {
  service_account_id = google_service_account.service_account.name
  role               = "roles/artifactregistry.reader"
  members = ["gcp-mig-demo@gcp-mig-demo-340020.iam.gserviceaccount.com"]

}