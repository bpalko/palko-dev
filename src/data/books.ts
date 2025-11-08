export interface Book {
  title: string;
  author: string;
  year: number;
  status: 'reading' | 'completed' | 'want-to-read';
  rating?: number; // out of 5
  notes?: string;
  coverUrl?: string; // Book cover image URL
  searchQuery?: string; // Custom search query for Google Books API (overrides title+author)
}

export const books: Book[] = [
  {
    title: 'Kitchen Confidential',
    author: 'Anthony Bourdain',
    year: 2000,
    status: 'completed',
    notes: 'My favorite writer of all time, if I\'m being honest. Everything he writes is visceral and raw. I\'m always drawn into his world and perspective on life.',
  },
  {
    title: 'Halo: The Fall of Reach',
    author: 'Eric Nylund',
    year: 2001,
    status: 'completed',
    notes: 'I love the Halo universe; I grew up in it. This is the first book in the "Halo Series". Highly recommend reading this to gain some lore on the Master Chief and how he came to be.',
  },
  {
    title: 'Halo: The Flood',
    author: 'William C. Dietz',
    year: 2003,
    status: 'reading',
    notes: 'Continuation of the Halo books, albeit the only one based around the first game, Halo: Combat Evolved. This has been a treat to read as I sunk many hours into the video game.',
  },
  {
    title: 'The Intelligent Investor',
    author: 'Benjamin Graham',
    year: 1949,
    status: 'completed',
    notes: 'One of my first introductions into financial literacy. A dense read. I probably grasped half of the subjects... probably should read it once more.',
  },
  {
    title: 'The Courage to Be Disliked',
    author: 'Ichiro Kishimi and Fumitake Koga',
    year: 2018,
    status: 'reading',
    notes: 'Reading this one off and on. Adlerian psychology and Stoic themes. Compelling.',
    searchQuery: 'The Courage to Be Disliked',
  },
  {
    title: 'The Bobiverse',
    author: 'Dennis E. Taylor',
    year: 2016,
    status: 'completed',
    notes: 'A series about a programmer who copies himself across the universe. Please — just read them...',
    searchQuery: 'For We Are Many Dennis Taylor',
  },
  {
    title: 'As A Man Thinketh',
    author: 'James Allen',
    year: 1903,
    status: 'completed',
    notes: 'One of those books you don\'t just read and forget about; you\'ll want to read this over and over. I sat and stared at each page for extended periods trying to grasp what Allen is getting across. The ideas have endured the test of time. Give it a shot.',
  },
  {
    title: 'Going Infinite: The Rise and Fall of a New Tycoon',
    author: 'Michael Lewis',
    year: 2023,
    status: 'completed',
    notes: 'I\'ve always had an interest in finances and crypto. SBF is an incredibly interesting person, no matter how you view the events of FTX or his life. Lewis made reading this seem like a movie.'
  },
  {
    title: 'Meditations',
    author: 'Marcus Aurelius, translated by Gregory Hays',
    year: 2003,
    status: 'completed',
    notes: 'It\'s Meditations. Read it a 100 times. I like the application of the Stoic mindset.'
  },
  {
    title: 'The Phoenix Project',
    author: 'Gene Kim, Kevin Behr, George Spafford',
    year: 2014,
    status: 'completed',
    notes: 'What software engineer hasn\'t read this? Well they should...'
  },
  {
    title: 'The Martian',
    author: 'Andy Weir',
    year: 2011,
    status: 'want-to-read',
    notes: 'I\'ve seen the movie many times; I heard the book makes Mark Watney\'s struggles even more intense.',
  },
  {
    title: 'The Big Short',
    author: 'Michael Lewis',
    year: 2010,
    status: 'want-to-read',
    notes: 'I\'ve seen the movie many times; need to read the book. I\'ve enjoyed Lewis\' writing style in the past.',
  },
  {
    title: 'Medium Raw: A Bloody Valentine to the World of Food and People Who Cook',
    author: 'Anthony Bourdain',
    year: 2010,
    status: 'want-to-read',
    notes: 'A follow-up to Kitchen Confidential. I\'m on a mission to read all of his books. I\'ll probably start back up here.',
  }
];
