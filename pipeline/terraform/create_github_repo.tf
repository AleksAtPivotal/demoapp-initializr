provider "github" {
  token        = "${var.github_token}" 
}

resource "github_repository" "githubrepo" {
  name        = "${var.reponame}"
  description = "Created by DemoApp-Initializr"

  private = false
  auto_init = true
}

variable "github_token" {
  type = "string"
}

variable "reponame" {
  type = "string"
}

