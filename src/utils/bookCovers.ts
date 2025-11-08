export interface GoogleBooksResponse {
  items?: Array<{
    volumeInfo: {
      imageLinks?: {
        thumbnail?: string;
        smallThumbnail?: string;
      };
    };
  }>;
}

/**
 * Fetches book cover image URL from Google Books API
 * @param title - Book title
 * @param author - Book author
 * @param customQuery - Optional custom search query (overrides title+author)
 * @returns URL to book cover image or null if not found
 */
export async function getBookCover(
  title: string,
  author: string,
  customQuery?: string
): Promise<string | null> {
  try {
    // Use custom query if provided, otherwise construct from title+author
    const query = customQuery
      ? encodeURIComponent(customQuery)
      : `intitle:${encodeURIComponent(title)}+inauthor:${encodeURIComponent(author)}`;
    const url = `https://www.googleapis.com/books/v1/volumes?q=${query}&maxResults=1`;

    const response = await fetch(url);
    const data: GoogleBooksResponse = await response.json();

    // Get the thumbnail URL, prefer larger thumbnail over small
    const imageLinks = data.items?.[0]?.volumeInfo?.imageLinks;
    const coverUrl = imageLinks?.thumbnail || imageLinks?.smallThumbnail;

    // Google Books returns http URLs, convert to https for security
    return coverUrl ? coverUrl.replace('http://', 'https://') : null;
  } catch (error) {
    console.error(`Failed to fetch cover for "${title}":`, error);
    return null;
  }
}

/**
 * Fetches cover images for multiple books in parallel
 * @param books - Array of books with title, author, and optional searchQuery
 * @returns Map of book titles to cover URLs
 */
export async function getBookCovers(
  books: Array<{ title: string; author: string; searchQuery?: string }>
): Promise<Map<string, string | null>> {
  const coverPromises = books.map(async (book) => {
    const coverUrl = await getBookCover(book.title, book.author, book.searchQuery);
    return [book.title, coverUrl] as const;
  });

  const results = await Promise.all(coverPromises);
  return new Map(results);
}
