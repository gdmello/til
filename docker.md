Delete all stopped containers
=============================

  docker rm $(docker ps -aq)
  
Delete all untagged images
==========================

  docker rmi $(docker images | grep "^<none>" | awk '{print $3}')
  