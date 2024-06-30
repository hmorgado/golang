variable "name" {
  type = string
}

variable "my_map" {
  type = map(any)
}

resource "local_file" "foo" {
  content  = "Hello, ${var.name}"
  filename = "${path.module}/hello.txt"
}