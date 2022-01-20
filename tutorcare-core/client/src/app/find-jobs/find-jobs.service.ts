import { Injectable } from '@angular/core';
import {Http} from '@angular/http';
import {environment} from '../../environments/environment';
import { Observable } from 'rxjs/Observable';
import { map } from 'rxjs/operators';
import { GeolocationPositionWithUser } from '../models/geolocationposition.model';
import { Post } from '../models/post.model';


@Injectable()
export class FindJobsService {
  results:Object[];
  _posts: Post[] | undefined;
  constructor(private http: Http) {
    this.results = []
  }

  getPosts(): Observable<Post[]> {
    let url = `/api/posts/active`;
    return new Observable((observer: any) => {
       this.http.get(url)
           .pipe(map((res: any) => res.json().posts))
           .subscribe((data: Post[]) => {
              this._posts = data

              observer.next(this._posts);
              observer.complete();
           });
    });
 }

//  getPosition(): Observable<any> {
//   return new Observable(observer => {
//     window.navigator.geolocation.getCurrentPosition(position => {
//       observer.next(position);
//       observer.complete();
//     },
//       error => observer.error(error));
//   });
// }

}
