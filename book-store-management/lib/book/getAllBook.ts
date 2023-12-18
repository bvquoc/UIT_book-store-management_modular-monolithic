import { apiKey, endPoint } from "@/constants";

export default async function getAllBooks(page: number) {
  const res = await fetch(`${endPoint}/v1/booktitles?page=${page}`, {
    headers: {
      accept: "application/json",
      Authorization: apiKey,
    },
    next: {
      revalidate: 0,
    },
  });
  if (!res.ok) {
    throw new Error("Failed to fetch data");
  }
  return res.json().then((json) => {
    return {
      paging: json.paging,
      data: json.data,
    };
  });
}
