import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-ranking',
  templateUrl: './ranking.component.html',
  styleUrls: ['./ranking.component.scss']
})
export class RankingComponent implements OnInit {

  public rankings = [
    {nickname:"Dario Ciaudano", points:"100"},
    {nickname:"Fabio Carfì", points:"90"},
    {nickname:"Amerigo Mancino", points:"70"},
    {nickname:"Simone Tiburzi", points:"60"},
    {nickname:"Fabiana Bonfante", points:"54"},
    {nickname:"Alessio Saccardo", points:"49"},
    {nickname:"Gennaro Lo Stupido", points:"20"},
  ]
  constructor() { }

  ngOnInit(): void {

  }

}
