export default async function UserPage({
  params,
}: {
  params: Promise<{ user_id: string }>;
}) {
  const user_id = (await params).user_id;
  return <p>{user_id}</p>;
}
