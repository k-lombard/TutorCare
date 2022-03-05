import { List } from "postcss/lib/list";
import { Application } from "./application.model";
import { User } from "./user.model";

export interface Post {
  user_id?: string
  caregiver_id?: string
  post_id?: number
  title?: string
  care_description?: string
  tags?: string
  tagList?: string[]
  care_type?: string
  completed?: boolean
  date_of_job?: string
  start_date?: string
  start_time?: string
  end_date?: string
  end_time?: string
  date_posted?: string
  applications?: Application[]
  caregiver?: User
  application_id?: number
  user?: User
}
