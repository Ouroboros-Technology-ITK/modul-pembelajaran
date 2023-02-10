import { useState } from "react";
import {
  Box,
  Image,
  Badge,
  SimpleGrid,
  Modal,
  ModalOverlay,
  ModalContent,
  ModalHeader,
  Button,
  ModalCloseButton,
  ModalBody,
  ModalFooter,
  useDisclosure,
} from "@chakra-ui/react";
import { StarIcon } from "@chakra-ui/icons";

function Task1(props) {
  const { property } = props;

  // TODO: answer here
}

function Task2(props) {
  const { properties } = props;

  // TODO: answer here
}

function Task3(props) {
  // TODO: answer here
  const { body, title } = props;
  return (
    <>
      <Button>Open Modal</Button>  {/* TODO: replace this */}

      {/* TODO: answer here */}
    </>
  );
}

function Task1_1(props) {
  return <div>{props.name}</div>;
}

function Task1_2() {
  const [name, setName] = useState("Ganti namaku");

  return (
    <div>
      <h1>{name}</h1>
      <button onClick={() => setName("Imam Assidiqqi")}>Change Name</button>
    </div>
  );
}

export { Task1, Task1_1, Task1_2, Task2, Task3 };
