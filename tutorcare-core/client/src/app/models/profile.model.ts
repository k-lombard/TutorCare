import { Review } from "./review.model";

export interface Profile {
    user_id       ?:string
    profile_id    ?:number
    profile_pic   ?:boolean
    bio           ?:string
    badge_list    ?:string
    age           ?:number
    gender        ?:string
    language      ?:string
    experience    ?:string
    education     ?:string
    skills        ?:string
    service_types ?:string
    age_groups    ?:string
    covid19       ?:boolean
    smoker        ?:boolean
    jobs_completed?:number
    rate_range    ?:string
    rating        ?:number
}

/*interface FamilyMember {
    name?: string
    DoB?: string
    gender?: string
    bio?: string
    hasMedication?: boolean
    medicationInfo?: string
    hasAllergy?: boolean
    allergyInfo?: string
}*/