import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-materials',
  templateUrl: './materials.component.html',
  styleUrls: ['./materials.component.css']
})
export class MaterialsComponent {
  @Input() materials: any[];
  
  constructor() {
    this.materials = [{name: 'Wilson US Open Tennis Ball',
    type: 'Tennis Ball',
    description: 'The Wilson US Open Tennis Ball is the official ball of the US Open since 1979.',
    image: 'https://dks.scene7.com/is/image/GolfGalaxy/WRT107_NOCOLOR_OPKG?qlt=70&wid=700&fmt=pjpeg&op_sharpen=1',
    price: '10.99$'
  },
  {
    name: 'Babolat Pure Aero 2019',
    type: 'Tennis Racquet',
    description: 'Babolat Pure Aero 2019 is designed for power and spin.',
    image: 'https://img.tennis-warehouse.com/watermark/rs.php?path=BPAR-1.jpg&nw=455',
    price: '20.99$'
    
  },
  {
    name: 'Head Hawk Touch Tennis String',
    type: 'Tennis String',
    description: 'The Head Hawk Touch string is designed for versatile players who want control and durability.',
    image: 'https://servingaces.com.au/wp-content/uploads/2020/05/SA-HHT-1.jpg',
    price: '15.99$'
  },
  {
    name: 'NikeCourt Air Zoom Vapor X',
    type: 'Tennis Shoes',
    description: 'The NikeCourt Air Zoom Vapor X provides ultimate control on the court.',
    image: 'https://www.tradeinn.com/f/13748/137480111/nike-chaussures-surface-dure-court-air-zoom-vapor-x.jpg',
    price: '30.99$'
  },
  {
    name: 'Roland Garros Clay',
    type: 'Court Surface',
    description: 'The Roland Garros clay court gives a high bounce and slow play, intensifying the rallies.',
    image: 'https://a.espncdn.com/photo/2015/0522/espnw_u_roalnd-garros_mb_1296x729.jpg',
    price: '18.99$'
  },
  {
    name: 'Dunlop Australian Open Tennis Ball',
    type: 'Tennis Ball',
    description: 'Dunlop Australian Open is the official ball of the Australian Open.',
    image: 'https://m.media-amazon.com/images/I/81wgGda8RiL._AC_UF1000,1000_QL80_.jpg',
    price: '60.99$'
  },
  {
    name: 'Yonex EZONE 98',
    type: 'Tennis Racquet',
    description: 'Yonex EZONE 98 offers a great balance of power, control, and comfort.',
    image: 'https://img.tennis-warehouse.com/watermark/rs.php?path=EZO98-1.jpg&nw=455',
    price: '10.99$'
  },
  {
    name: 'Luxilon ALU Power Rough Tennis String',
    type: 'Tennis String',
    description: 'Luxilon ALU Power Rough provides a combination of power and spin for aggressive players.',
    image: 'https://www.luxilon.com/en-us/media/catalog/product/W/R/WRZ995200__9b8cb239ba40562d2957d237c96709de.png',
    price: '100.99$'
  },
  {
    name: 'Adidas SoleCourt Boost',
    type: 'Tennis Shoes',
    description: 'The Adidas SoleCourt Boost offers comfort and stability on the court.',
    image: 'https://assets.adidas.com/images/w_600,f_auto,q_auto/b1e16f607ea34694893fa98400f224d9_9366/SoleCourt_Shoes_Black_AH2131_01_standard.jpg',
    price: '60.99$'
  },
  {
    name: 'Wimbledon Grass',
    type: 'Court Surface',
    description: 'The Wimbledon grass court offers a fast game and low bounce, favoring serve-and-volley players.',
    image: 'https://photo-assets.wimbledon.com/images/pics/large/s_grounds_4757_29062019_tc.jpg',
    price: '50.99$'
  }];
  }
  toggleDescription(material: any) {
    material.showDescription = !material.showDescription;
  }

}
