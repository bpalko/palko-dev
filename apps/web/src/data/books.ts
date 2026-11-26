export interface Book {
  title: string;
  author: string;
  year: number;
  status: 'reading' | 'completed' | 'want-to-read';
  rating?: number; // out of 5
  notes?: string;
  coverUrl?: string; // Book cover image URL
  searchQuery?: string; // Custom search query for the Open Library API (overrides title+author)
}

export const books: Book[] = [
  {
    title: 'Kitchen Confidential',
    author: 'Anthony Bourdain',
    year: 2000,
    status: 'completed',
    notes: 'My favorite writer of all time, if I\'m being honest. Everything he writes is visceral and raw. I\'m always drawn into his world and perspective on life.',
    coverUrl: 'https://covers.openlibrary.org/b/id/843879-L.jpg',
  },
  {
    title: 'Halo: The Fall of Reach',
    author: 'Eric Nylund',
    year: 2001,
    status: 'completed',
    notes: 'I love the Halo universe; I grew up in it. This is the first book in the "Halo Series". Highly recommend reading this to gain some lore on the Master Chief and how he came to be.',
    coverUrl: 'https://covers.openlibrary.org/b/id/1726374-L.jpg',
  },
  {
    title: 'Halo: The Flood',
    author: 'William C. Dietz',
    year: 2003,
    status: 'completed',
    notes: 'Continuation of the Halo books, albeit the only one based around the first game, Halo: Combat Evolved. This has been a treat to read as I sunk many hours into the video game.',
    coverUrl: 'https://covers.openlibrary.org/b/id/210887-L.jpg',
  },
  {
    title: 'The Intelligent Investor',
    author: 'Benjamin Graham',
    year: 1949,
    status: 'completed',
    notes: 'One of my first introductions into financial literacy. A dense read. I probably grasped half of the subjects... probably should read it once more.',
    coverUrl: 'https://covers.openlibrary.org/b/id/10826961-L.jpg',
  },
  {
    title: 'The Courage to Be Disliked',
    author: 'Ichiro Kishimi and Fumitake Koga',
    year: 2018,
    status: 'reading',
    notes: 'Reading this one off and on. Adlerian psychology and Stoic themes. Compelling.',
    coverUrl: 'https://covers.openlibrary.org/b/id/10873626-L.jpg',
  },
  {
    title: 'The Bobiverse',
    author: 'Dennis E. Taylor',
    year: 2016,
    status: 'completed',
    notes: 'A series about a programmer who copies himself across the universe. Please — just read them...',
    coverUrl: 'https://covers.openlibrary.org/b/id/12750192-L.jpg',
  },
  {
    title: 'As A Man Thinketh',
    author: 'James Allen',
    year: 1903,
    status: 'completed',
    notes: 'One of those books you don\'t just read and forget about; you\'ll want to read this over and over. I sat and stared at each page for extended periods trying to grasp what Allen is getting across. The ideas have endured the test of time. Give it a shot.',
    coverUrl: 'https://covers.openlibrary.org/b/id/6268048-L.jpg',
  },
  {
    title: 'Going Infinite: The Rise and Fall of a New Tycoon',
    author: 'Michael Lewis',
    year: 2023,
    status: 'completed',
    notes: 'I\'ve always had an interest in finances and crypto. SBF is an incredibly interesting person, no matter how you view the events of FTX or his life. Lewis made reading this seem like a movie.',
    coverUrl: 'https://covers.openlibrary.org/b/id/14559785-L.jpg',
  },
  {
    title: 'Meditations',
    author: 'Marcus Aurelius, translated by Gregory Hays',
    year: 2003,
    status: 'completed',
    notes: 'It\'s Meditations. Read it a 100 times. I like the application of the Stoic mindset.',
    coverUrl: 'https://covers.openlibrary.org/b/id/211529-L.jpg',
  },
  {
    title: 'The Phoenix Project',
    author: 'Gene Kim, Kevin Behr, George Spafford',
    year: 2014,
    status: 'completed',
    notes: 'What software engineer hasn\'t read this? Well they should...',
    coverUrl: 'https://covers.openlibrary.org/b/id/9151976-L.jpg',
  },
  {
    title: 'The Martian',
    author: 'Andy Weir',
    year: 2011,
    status: 'completed',
    notes: 'I\'ve seen the movie many times; I heard the book makes Mark Watney\'s struggles even more intense.',
    coverUrl: 'https://covers.openlibrary.org/b/id/11447888-L.jpg',
  },
  {
    title: 'Dungeon Crawler Carl',
    author: 'Matt Dinniman',
    year: 2020,
    status: 'reading',
    notes: 'A LitRPG where a man and his ex-girlfriend\'s cat survive a deadly alien game show built inside a dungeon. Ridiculous premise, hard to put down.',
    coverUrl: 'https://covers.openlibrary.org/b/id/15143022-L.jpg',
  },
  {
    title: 'Red Rising',
    author: 'Pierce Brown',
    year: 2014,
    status: 'reading',
    notes: 'Sci-fi caste system, rebellion, and a lot of betrayal. Everyone told me to read this one for years.',
    coverUrl: 'https://covers.openlibrary.org/b/id/14651895-L.jpg',
  },
  {
    title: 'The Big Short',
    author: 'Michael Lewis',
    year: 2010,
    status: 'want-to-read',
    notes: 'I\'ve seen the movie many times; need to read the book. I\'ve enjoyed Lewis\' writing style in the past.',
    coverUrl: 'https://covers.openlibrary.org/b/id/6386926-L.jpg',
  },
  {
    title: 'Medium Raw: A Bloody Valentine to the World of Food and People Who Cook',
    author: 'Anthony Bourdain',
    year: 2010,
    status: 'want-to-read',
    notes: 'A follow-up to Kitchen Confidential. I\'m on a mission to read all of his books. I\'ll probably start back up here.',
    coverUrl: 'https://covers.openlibrary.org/b/id/6298208-L.jpg',
  }
];
