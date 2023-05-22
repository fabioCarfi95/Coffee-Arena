import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Location } from '@angular/common';
import { LOGIN } from 'src/app/constants/constants';

@Component({
  selector: 'app-toolbar',
  templateUrl: './toolbar.component.html',
  styleUrls: ['./toolbar.component.scss']
})
export class ToolbarComponent implements OnInit {

  public time: Date = new Date();
  public logged: boolean;

  private timer;

  constructor(private location: Location, private router: Router) {

    this.logged = false;

    this.timer = setInterval(() => {
      this.time = new Date();
    }, 1000);

    this.location.onUrlChange( () => {
      let url = this.location.path();
      let url_segments = url.slice(1,url.length).split("/");

      if (!url_segments.includes(LOGIN)){
        this.logged = true;
      }
    })
  }

  ngOnInit(): void { }

  logout() {
    this.logged = false;
    this.router.navigate([LOGIN]);
  }
}
