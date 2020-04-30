variable "docker_network" {
  type        = string
  description = "the name of the docker network"
  default     = "game"
}


variable "dbport" {
    type = map
    default = {
        "internal"  = 3000
        "external"  = 80
}
    }



variable "dbgame" {
  type        = string
  description = "the name of the container db"
  default     = "mysql:5.7"
}



variable "battlegame" {
  type        = string
  description = "the name of the container api"
  default     = "fake-backend"
}
