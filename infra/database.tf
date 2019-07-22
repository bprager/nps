resource "helm_release" "customerDatabase" {
  name  = "customerDatabase"
  chart = "stable/postgresql"

  set {
    name  = "secret"
    value = "noIdea"
  }

  set {
    name  = "postgresqlUsername"
    value = "user"
  }

  set {
    name  = "postgresqlPassword"
    value = "droWssaP"
  }
}
