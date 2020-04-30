# Configuration ressource 

# create a network 
resource "docker_network" "docker_network" {
  name   = var.docker_network
  driver = "bridge" 
}

# Create a container
resource "docker_container" "dbgame" {
  image = var.dbgame
  name  = "dbgame"

 
  mounts {
    type   = "bind"
    source = "/var/"
    target = "/var/lib/mysql"
  }
  env = [" MYSQL_ROOT_PASSWORD: rootpwdgame"," MYSQL_DATABASE: battleboat"," MYSQL_USER: battleuser","MYSQL_PASSWORD: battlepass"]
  restart = "always" 

networks_advanced {
    name = var.docker_network
  }
}




#### Configuration of API 




resource "docker_container" "battlegame" {
  image = var.battlegame
  name  = "battlegame"

   ports {
    internal = var.dbport["internal"]
    external = var.dbport["external"]
  }
  mounts {
    type   = "bind"
    source = "/home/centos/projet-fake-backend-battleboat/terraform-fake-backend/battleboat/"
    target = "/etc/backend/static"
  }
  env = ["DATABASE_HOST: dbgame"," DATABASE_PORT: 3306","DATABASE_USER: battleuser"," DATABASE_PASSWORD: battlepass","DATABASE_NAME: battleboat"]
  restart = "always"

networks_advanced {
    name = var.docker_network
  }
}


