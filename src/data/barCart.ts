export interface Booze {
  name: string;
  type: string; // e.g., "Bourbon", "Scotch", "Rye", "Gin"
  distillery?: string;
  age?: string; // e.g., "12 Year", "18 Year"
  proof?: number;
  rating?: number; // out of 5
  notes?: string;
  imageUrl?: string; // Optional bottle image URL
  status: 'current' | 'finished' | 'wishlist';
}

export const barCart: Booze[] = [
  {
    name: 'Buffalo Trace',
    type: 'Bourbon',
    distillery: 'Buffalo Trace Distillery',
    proof: 90,
    rating: 4,
    notes: 'A solid everyday bourbon . You\'ll be hard pressed to find something better for the price.',
    status: 'current',
    imageUrl: '/bottles/buffalotrace.png',
  },
  {
    name: 'Midwinter Night\'s Dram - Act 13',
    type: 'Whiskey',
    distillery: 'Highwest Distillery',
    proof: 99,
    rating: 5,
    notes: 'I like this year. Warm, vanilla and spices. ',
    status: 'current',
    imageUrl: '/bottles/midwinter13.png',
  },
  {
    name: 'Midwinter Night\'s Dram - Act 12',
    type: 'Whiskey',
    distillery: 'Highwest Distillery',
    proof: 98.6,
    rating: 3,
    notes: 'Kinda average compared to Act 11 and Act 13.',
    status: 'current',
    imageUrl: '/bottles/midwinter12.png',
  },
  {
    name: 'EH Taylor Small Batch',
    type: 'Bourbon',
    distillery: 'Buffalo Trace Distillery',
    proof: 100,
    rating: 4,
    notes: 'It\s a solid bourbon. Picked it up while in Kentucky visiting the Buffalo Trace distillery',
    status: 'current',
    imageUrl: '/bottles/ehtaylorSB.png',
  },
  {
    name: 'Eagle Rare 10 Year',
    type: 'Bourbon',
    distillery: 'Buffalo Trace Distillery',
    proof: 90,
    rating: 4,
    notes: 'It\s a solid bourbon. Can\t go wrong with Eagle Rare. I have a few bottles at this rate.',
    status: 'current',
    imageUrl: '/bottles/eaglerare.png',
  },
  {
    name: 'Elijah Craig Barrel Proof - Store Pick Austin TX',
    type: 'Bourbon',
    distillery: 'Heaven Hill Distillery',
    proof: 129,
    rating: 3,
    notes: 'Picked it up on the way out of Austin TX. A bit too hot for my taste, but decent flavor underneath. The 129 isn\'t really the issue; I have higher proofs that drink easier.',
    status: 'current',
    imageUrl: '/bottles/elijahcraigBP.png',
  },
  {
    name: 'Tenjaku',
    type: 'Gin',
    distillery: 'Tenjaku',
    proof: 86,
    rating: 4,
    notes: 'Been enjoying this gin for a while now. Smooth with nice botanical flavors. Great for cocktails or sipping neat.',
    status: 'current',
    imageUrl: '/bottles/tenjaku.png',
  },
  {
    name: 'Campari',
    type: 'Aperitif',
    distillery: 'Campari Group',
    proof: 48,
    rating: 4,
    notes: 'Negroni. That is all.',
    status: 'current',
    imageUrl: '/bottles/campari.png',
  },
  {
    name: 'Rosso',
    type: 'Vermouth',
    distillery: 'Martini & Rossi',
    proof: 30,
    rating: 4,
    notes: 'Negroni. Cheap and available sweet vermouth.',
    status: 'current',
    imageUrl: '/bottles/rosso.png',
  },
];
