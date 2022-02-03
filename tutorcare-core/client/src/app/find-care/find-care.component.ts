import {Component, OnInit, ChangeDetectionStrategy} from '@angular/core';
import { Router } from '@angular/router';
import { select, Store } from '@ngrx/store';
import { getCurrUser } from '../auth/auth.selectors';
import { ChatroomsService } from '../find-jobs/chatrooms/chatrooms.service';
import { Chatroom } from '../models/chatroom.model';
import { GeolocationPositionWithUser } from '../models/geolocationposition.model';
import { User } from '../models/user.model';
import { AppState } from '../reducers';
import { FindCareService } from './find-care.service';

interface FilterOption {
    value: string;
    viewValue: string;
  }
@Component({
  selector: 'find-care-component',
  templateUrl: './find-care.component.html',
  styleUrls: ['./find-care.component.scss']
})
export class FindCareComponent implements OnInit {
    user!: User
    userId!: string
    userType!: string
    markers: any[] | undefined;
    start_index: number = 0
    zoom = 12
    center!: google.maps.LatLngLiteral
    options: google.maps.MapOptions = {
        zoomControl: false,
        scrollwheel: false,
        disableDoubleClickZoom: true,
        maxZoom: 15,
        minZoom: 8,
        styles: [
            {
                "featureType": "administrative",
                "elementType": "labels.text.fill",
                "stylers": [
                    {
                        "color": "#444444"
                    }
                ]
            },
            {
                "featureType": "landscape",
                "elementType": "all",
                "stylers": [
                    {
                        "color": "#f2f2f2"
                    }
                ]
            },
            {
                "featureType": "landscape",
                "elementType": "geometry.fill",
                "stylers": [
                    {
                        "visibility": "on"
                    },
                    {
                        "hue": "#ff0000"
                    }
                ]
            },
            {
                "featureType": "landscape.man_made",
                "elementType": "geometry",
                "stylers": [
                    {
                        "lightness": "100"
                    }
                ]
            },
            {
                "featureType": "landscape.man_made",
                "elementType": "labels",
                "stylers": [
                    {
                        "visibility": "off"
                    }
                ]
            },
            {
                "featureType": "landscape.natural",
                "elementType": "geometry.fill",
                "stylers": [
                    {
                        "lightness": "100"
                    }
                ]
            },
            {
                "featureType": "landscape.natural",
                "elementType": "labels",
                "stylers": [
                    {
                        "visibility": "off"
                    }
                ]
            },
            {
                "featureType": "landscape.natural.landcover",
                "elementType": "geometry.fill",
                "stylers": [
                    {
                        "visibility": "on"
                    }
                ]
            },
            {
                "featureType": "landscape.natural.terrain",
                "elementType": "geometry",
                "stylers": [
                    {
                        "lightness": "100"
                    }
                ]
            },
            {
                "featureType": "landscape.natural.terrain",
                "elementType": "geometry.fill",
                "stylers": [
                    {
                        "visibility": "off"
                    },
                    {
                        "lightness": "23"
                    }
                ]
            },
            {
                "featureType": "poi",
                "elementType": "all",
                "stylers": [
                    {
                        "visibility": "off"
                    }
                ]
            },
            {
                "featureType": "road",
                "elementType": "all",
                "stylers": [
                    {
                        "saturation": -100
                    },
                    {
                        "lightness": 45
                    }
                ]
            },
            {
                "featureType": "road.highway",
                "elementType": "all",
                "stylers": [
                    {
                        "visibility": "simplified"
                    }
                ]
            },
            {
                "featureType": "road.highway",
                "elementType": "geometry.fill",
                "stylers": [
                    {
                        "color": "#ffd900"
                    }
                ]
            },
            {
                "featureType": "road.arterial",
                "elementType": "labels.icon",
                "stylers": [
                    {
                        "visibility": "off"
                    }
                ]
            },
            {
                "featureType": "transit",
                "elementType": "all",
                "stylers": [
                    {
                        "visibility": "off"
                    }
                ]
            },
            {
                "featureType": "water",
                "elementType": "all",
                "stylers": [
                    {
                        "color": "#ffd900"
                    },
                    {
                        "visibility": "on"
                    }
                ]
            },
            {
                "featureType": "water",
                "elementType": "geometry.fill",
                "stylers": [
                    {
                        "visibility": "on"
                    },
                    {
                        "color": "#cccccc"
                    }
                ]
            },
            {
                "featureType": "water",
                "elementType": "labels",
                "stylers": [
                    {
                        "visibility": "off"
                    }
                ]
            }
        ]
    }

    selectedValue: string | undefined
    userCategory: string = ""
    filter_options: FilterOption[] = [
        {value: 'rating-0', viewValue: 'Rating Asc.'},
        {value: 'rating-1', viewValue: 'Rating Desc.'},
        {value: 'jobs-2', viewValue: 'Jobs Asc.'},
        {value: 'jobs-3', viewValue: 'Jobs Desc.'},
        {value: 'first-4', viewValue: 'Most Recent'},
        {value: 'first-5', viewValue: 'Least Recent'}
    ];

    locs!: GeolocationPositionWithUser[]
    constructor(private router: Router, private findCare: FindCareService, private chatroomService: ChatroomsService, private store: Store<AppState>) {}

    ngOnInit() {
        navigator.geolocation.getCurrentPosition((position) => {
            this.center = {
                lat: position.coords.latitude,
                lng: position.coords.longitude,
            }
        })
        this.findCare.getLocs().subscribe(data => {
            this.locs = data
            console.log(this.locs)
            let marks: any[] = []
            if (this.locs) {
              for(const loc of this.locs) {
                  console.log(loc)
                  var marker = JSON.parse(JSON.stringify({
                      "position": {
                          "lat": loc.latitude,
                          "lng": loc.longitude
                      }
                  }))
                  console.log(marker)
                  marks.push(marker)
                  console.log(marks)
              }
              this.markers = marks
              console.log(this.markers)
          }
        })
        this.store
        .pipe(
            select(getCurrUser)
        ).subscribe(data =>  {
            this.user = data
            this.userId = this.user.user_id || ""
            this.userType = this.user.user_category
      })

    }
    onFindCareClick() {
        this.router.navigate(['/find-care'])
    }

    onMessageClick(userid1: string) {
      this.chatroomService.getChatroomByTwoUsers(userid1, this.userId).subscribe((chatroom: Chatroom) => {
        this.router.navigate([`/find-jobs/messages/${chatroom.chatroom_id}`])
        this.chatroomService.setSelected(chatroom.chatroom_id)
      })
    }



}
