import {
  Button,
  Card,
  CardBody,
  CardFooter,
  CardHeader,
  Checkbox,
  Heading,
  Link,
  Stack,
  Text,
} from "@chakra-ui/react"
import { Link as RouterLink } from "react-router-dom"
import { useWorkspace } from "../../contexts"
import { Routes } from "../../routes"
import { TWorkspaceID } from "../../types"

type TWorkspaceCardProps = Readonly<{
  workspaceID: TWorkspaceID
  onSelectionChange?: (isSelected: boolean) => void
}>

export function WorkspaceCard({ workspaceID, onSelectionChange }: TWorkspaceCardProps) {
  const workspace = useWorkspace(workspaceID)

  if (workspace.data === undefined) {
    return null
  }

  const { id, provider, status, ide } = workspace.data

  return (
    <Card key={id} direction={{ base: "column", sm: "row" }} overflow="hidden" variant="outline">
      <Stack>
        <CardHeader display="flex" width="full" justifyContent="space-between">
          <Heading size="md">
            <Link as={RouterLink} to={Routes.toWorkspace(id)}>
              <Text fontWeight="bold">{id}</Text>
            </Link>
          </Heading>
          {onSelectionChange !== undefined && (
            <Checkbox onChange={(e) => onSelectionChange(e.target.checked)} />
          )}
        </CardHeader>
        <CardBody>
          {provider?.name && <Text>Provider: {provider.name}</Text>}
          <Text>Status: {status}</Text>
        </CardBody>

        <CardFooter>
          <Button
            colorScheme="primary"
            onClick={() => workspace.start({ ideConfig: ide })}
            isLoading={workspace.current?.name === "start"}>
            Start
          </Button>
          <Button onClick={() => workspace.stop()} isLoading={workspace.current?.name === "stop"}>
            Stop
          </Button>
          <Button
            colorScheme="red"
            onClick={() => workspace.remove()}
            isLoading={workspace.current?.name === "remove"}>
            Delete
          </Button>
        </CardFooter>
      </Stack>
    </Card>
  )
}