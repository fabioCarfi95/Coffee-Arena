import { Component, OnInit } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Router } from '@angular/router'

import { QUIZ_ROOM, SERVER_HOST } from 'src/app/constants/constants';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss']
})
export class HomeComponent implements OnInit {

  public sessionStart: boolean;
  public sessionAvailableTitleMessage: string;
  public sessionAvailableTimer: string;
  public game_choices = [
    {title: "1 vs 1", description: "Give your best versus a college", background: "", room: "1"},
    {title: "2 vs 2", description: "Lets create a team to win", background: "", room: "2"},
    {title: "Arena",  description: "One against all", background: "", room: "3"}
  ];

  constructor(private router: Router, private http: HttpClient) {
    this.sessionStart = false;
    this.sessionAvailableTitleMessage = "Next session will start in";
    this.sessionAvailableTimer = "";
  }

  ngOnInit(): void {
    let httpAccessOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json'
      })
    };
    this.http.get<any>(SERVER_HOST+"nextsession", httpAccessOptions).subscribe(
      (data) => {
        let next_session = data["next-session"];
        let session_datetime = new Date(parseInt(next_session)).valueOf();

        let timer = setInterval(() => {
          let date_now = Math.floor(new Date(Date.now()).valueOf() / 1000);
          let diff_in_seconds = (session_datetime - date_now);
          let hours = Math.floor(diff_in_seconds/3600);
          let minutes = Math.floor((diff_in_seconds%3600)/60);
          let seconds = (diff_in_seconds % 60);
          this.sessionAvailableTimer = hours.toString().padStart(2,'0')+"h:"+
                                       minutes.toString().padStart(2,'0')+"m:"+
                                       seconds.toString().padStart(2,'0')+"s";
          if (hours==0 && minutes==0 && seconds==0){
            this.sessionStart = true;
            clearInterval(timer);
          }
        }, 1000);

      },
      (error) => {
        console.log(error)
      }
    );
  }

  join(room: string){
    let type_room = parseInt(room);
    this.router.navigate([QUIZ_ROOM], {queryParams: {room: type_room}});
  }

}
