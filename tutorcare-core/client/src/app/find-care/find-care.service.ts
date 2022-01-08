import { Injectable } from '@angular/core';
import {Http} from '@angular/http';
import {environment} from '../../environments/environment';
import { Observable } from 'rxjs/Observable';
import { map } from 'rxjs/operators';
import { GeolocationPositionWithUser } from '../models/geolocationposition.model';


@Injectable() 
export class FindCareService {
  results:Object[];
  _positions: GeolocationPositionWithUser[] | undefined;
  constructor(private http: Http) { 
    this.results = []
  }

  getLocs(): Observable<GeolocationPositionWithUser[]> {
    let url = `/api/geolocationpositions/caregivers`;
    return new Observable((observer: any) => {
       this.http.get(url)
           .pipe(map((res: any) => res.json().geolocation_positions))
           .subscribe((data: GeolocationPositionWithUser[]) => {
              this._positions = data
 
              observer.next(this._positions);
              observer.complete();
           });
    });
 }

 getPosition(): Observable<any> {
  return new Observable(observer => {
    window.navigator.geolocation.getCurrentPosition(position => {
      observer.next(position);
      observer.complete();
    },
      error => observer.error(error));
  });
}

}