import React, { useState, useEffect, SetStateAction, Dispatch } from "react";
import {
  View,
  Text,
  StyleSheet,
  Dimensions,
  TouchableOpacity,
  Modal,
  FlatList,
  TouchableWithoutFeedback,
} from "react-native";
import { Colors } from "@/theme/variables";
import MaterialCommunityIcons from "react-native-vector-icons/MaterialCommunityIcons";
import DropDownPicker from "react-native-dropdown-picker";
import { rooms } from "@/services";

interface RoomsHeaderProps {
  onSelectRoom: (roomId: string) => void;
}

interface Room {
  room_name: string;
  room_id: string;
}

interface PopupProps {
  isVisible: boolean;
  onClose: () => void;
}

const Popup: React.FC<PopupProps> = ({
  isVisible,
  onClose,
}) => {
  const [value, setValue] = useState("");
  const handleClick = () => {
    onClose();
  };
  return (
    <Modal
      animationType="none"
      transparent={true}
      visible={isVisible}
      onRequestClose={onClose}
    >
      <TouchableWithoutFeedback onPress={onClose}>
        <View style={styles.modalContainer}>
          <View style={styles.modalContent}>
            <Text style={styles.label}>Contact Information</Text>
            <View style={styles.informContainer}>
              <Text style={styles.informText}>Room B - Nguyen Van D</Text>
            </View>
            <TouchableOpacity onPress={handleClick} style={styles.button}>
              <Text style={styles.buttontext}>Close</Text>
            </TouchableOpacity>
          </View>
        </View>
      </TouchableWithoutFeedback>
    </Modal>
  );
};

const screenWidth = Dimensions.get("screen").width;
const screenHeight = Dimensions.get("screen").height;

export const RoomsHeader: React.FC<RoomsHeaderProps> = ({ onSelectRoom }) => {
  const [open, setOpen] = useState(false);
  const [selectedValue, setSelectedValue] = useState<string>("");
  const [items, setItems] = useState<{ label: string; value: string }[]>([]);
  const [isModalVisible, setIsModalVisible] = useState(false);

  async function fetchData() {
    try {
      const roomData = await rooms.getroominform();
      const roomNames = roomData.data.map((room: Room) => ({
        label: room.room_name,
        value: room.room_id,
      }));
      setItems(roomNames);
    } catch (error) {
      console.error("Error fetching room data:", error);
    }
  }

  useEffect(() => {
    fetchData();
  }, []);

  useEffect(() => {
    if (items.length > 0 && selectedValue === "") {
      const defaultRoomId = items[0].value;
      setSelectedValue(defaultRoomId);
      onSelectRoom(defaultRoomId);
      // console.log(defaultRoomId);
    }
  }, [items, selectedValue, onSelectRoom]);

  const handleContactPress = () => {
    setIsModalVisible(true);
  };

  const handleDropdownChange = (value: SetStateAction<string>) => {
    const selectedRoomId = typeof value === 'function' ? (value as (prevState: string) => string)(selectedValue) : value;
    setSelectedValue(selectedRoomId);
    onSelectRoom(selectedRoomId);
    // console.log(selectedRoomId);
  };

  return (
    <View style={styles.headerContainer}>
      <DropDownPicker
        open={open}
        items={items}
        value={selectedValue}
        setOpen={setOpen}
        setValue={handleDropdownChange}
        setItems={setItems}
        placeholder="Rooms:"  
        style={{
          borderRadius: 10,
          borderWidth: 2.5,
          borderColor: "rgba(27, 97, 181, 0.89)",
        }}
        containerStyle={styles.dropdownContainer}
        dropDownContainerStyle={{
          borderTopWidth: 0,
          borderRadius: 10,
          borderWidth: 2.5,
          borderColor: "rgba(27, 97, 181, 0.89)",
        }}
      />

      <TouchableOpacity style={styles.iconContainer} onPress={handleContactPress}>
        <MaterialCommunityIcons
          name="account-box"
          size={50}
          color={"rgba(27, 97, 181, 0.89)"}
        />
      </TouchableOpacity>

      {isModalVisible && (
        <Popup
          isVisible={isModalVisible}
          onClose={() => setIsModalVisible(false)}
        />
      )}

    </View>
  );
};

const styles = StyleSheet.create({
  headerContainer: {
    flexDirection: "row",
    paddingHorizontal: 5,
    backgroundColor: Colors.WHITE,
    paddingBottom: 0.02 * screenHeight,
    shadowColor: "#000",
    shadowOffset: { width: 1, height: 1 },
    shadowOpacity: 0.4,
    shadowRadius: 3,
    elevation: 5,
    zIndex: 10,
  },
  dropdownContainer: {
    flex: 0.8,
    flexDirection: "row",
    marginLeft: 10,
    marginRight: 5,
    backgroundColor: Colors.WHITE,
  },
  iconContainer: {
    flex: 0.2,
    alignItems: "center",
    justifyContent: "center",
  },
  modalContainer: {
    flex: 1,
    justifyContent: "center",
    alignItems: "center",
    backgroundColor: "rgba(0,0,0,0.5)",
  },
  modalContent: {
    padding: 16,
    borderRadius: 10,
    elevation: 5,
    width: 0.8 * screenWidth,
    backgroundColor: Colors.WHITE,
  },
  informContainer: {
    padding: 10,
    marginVertical: 10,
    borderRadius: 10,
    alignItems: "center",
    backgroundColor: "rgba(27, 97, 181, 0.10)",
  },
  label: {
    fontSize: 20,
    fontFamily: "Poppins_400Regular",
    color: "rgba(27, 97, 181, 0.89)",
  },
  button: {
    justifyContent: "center",
    paddingHorizontal: 10,
    paddingVertical: 4,
    borderWidth: 1,
    borderRadius: 10,
    borderColor: "#1B61B5",
    marginTop: 5,
    backgroundColor: Colors.WHITE,
  },
  buttontext: {
    color: "#1B61B5",
    fontFamily: "Poppins_400Regular",
    textAlign: "center",
    textAlignVertical: "center",
  },
  informText: {
    fontSize: 16,
    color: "rgba(27, 97, 181, 0.89)",
    fontFamily: "Poppins_500Medium",
  },
});
