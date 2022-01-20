import {Component, OnInit, ChangeDetectionStrategy} from '@angular/core';
import { Router } from '@angular/router';
import { GeolocationPositionWithUser } from '../models/geolocationposition.model';
import { Post } from '../models/post.model';
import { FindJobsService } from './find-jobs.service';

interface FilterOption {
    value: string;
    viewValue: string;
  }
@Component({
  selector: 'find-jobs-component',
  templateUrl: './find-jobs.component.html',
  styleUrls: ['./find-jobs.component.scss']
})
export class FindJobsComponent implements OnInit {
    selectedValue: string | undefined
    rate: number = 4.5
    userCategory: string = ""
    posts!: Post[]
    search: string =""
    filter_options: FilterOption[] = [
        {value: 'rating-0', viewValue: 'Type: Tutoring'},
        {value: 'rating-1', viewValue: 'Type: Babysitting'},
        {value: 'jobs-2', viewValue: 'Type: Other'}
    ];
    constructor(private router: Router, private findJobs: FindJobsService) {}

    ngOnInit() {
        this.findJobs.getPosts().subscribe(data => {
            this.posts = data
            console.log(this.posts)
        })

    }
    onFindCareClick() {
        this.router.navigate(['/find-care'])
    }



}
