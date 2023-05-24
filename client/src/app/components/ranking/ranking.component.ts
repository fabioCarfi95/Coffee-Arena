import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-ranking',
  templateUrl: './ranking.component.html',
  styleUrls: ['./ranking.component.scss']
})
export class RankingComponent implements OnInit {

  public rankings = [
    {nickname:"Dario", points:"100"},
    {nickname:"Fabio", points:"90"},
    {nickname:"Luca", points:"70"},
    {nickname:"Simone", points:"60"},
    {nickname:"Fabiana", points:"54"},
    {nickname:"Alessio", points:"49"},
    {nickname:"Gennaro", points:"20"},
  ]
  constructor() { }

  ngOnInit(): void {

  }

}
