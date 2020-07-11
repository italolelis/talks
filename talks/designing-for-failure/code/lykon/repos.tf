module "my-cool-srv" {
  source = "./modules/repo"

  name                   = "my-cool-srv"
  description            = "This is my cool service description"
  private                = true
  delete_branch_on_merge = true
  topics                 = ["go", "is", "the", "best"]
  gitignore_template     = "Go"
  auto_init              = true
  teams                  = { (github_team.developers.id) = "push" }
  secrets                = local.secrets
  branch_protection      = local.default_tests_branch_protection
  template               = { // HLtemplate
    owner      = "my-org" // HLtemplate
    repository = "go-srv-template" // HLtemplate
  }
}
