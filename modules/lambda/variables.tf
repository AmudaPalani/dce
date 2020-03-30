variable "namespace" {
  type = string
}
variable "environment" {
  type    = map(string)
  default = { TERRAFORM = "true" }
}
variable "global_tags" {
  type = map(string)
}
variable "name" {
  type = string
}
variable "description" {
  type = string
}
variable "timeout" {
  type    = number
  default = 300
}
variable "handler" {
  type = string
}
variable "alarm_topic_arn" {
  type        = string
  description = "ARN of SNS Topic, for alarm notifications"
}

variable "dlq_enabled" {
  type        = bool
  description = "dead letter config enabled"
  default     = false
}

variable "dead_letter_config" {
  type = object({
    target_arn = string
  })
  description = "target_arn for dlq"
  default = null
}
