import React, { useState, useEffect } from "react";
import {
  Dimensions,
  View,
  Text,
  StyleSheet,
  TouchableOpacity,
  TextInput,
  Modal,
  Button,
  TouchableWithoutFeedback,
} from "react-native";
import { BarCodeScanner } from "expo-barcode-scanner";
import BarcodeMask from "react-native-barcode-mask";
import { Colors } from "@/theme/variables";
import { Header } from "@/components/header";

const screenWidth = Dimensions.get("screen").width;
const screenHeight = Dimensions.get("screen").height;

interface ResultPopupProps {
  isVisible: boolean;
  result: string | null;
  onClose: () => void;
}

interface InputPopupProps {
  isVisible: boolean;
  onClose: () => void;
}

const ResultPopup: React.FC<ResultPopupProps> = ({
  isVisible,
  result,
  onClose,
}) => {
  return (
    <Modal
      animationType="fade"
      transparent={true}
      visible={isVisible}
      onRequestClose={onClose}
    >
      <TouchableWithoutFeedback onPress={onClose}>
        <View style={styles.modalContainer}>
          <View style={styles.modalContent}>
            <Text>Result: {result}</Text>
            <Button color={Colors.PRIMARY} title="CLOSE" onPress={onClose} />
          </View>
        </View>
      </TouchableWithoutFeedback>
    </Modal>
  );
};

const InputPopup: React.FC<InputPopupProps> = ({ isVisible, onClose }) => {
  return (
    <Modal
      animationType="fade"
      transparent={true}
      visible={isVisible}
      onRequestClose={onClose}
    >
      <TouchableWithoutFeedback onPress={onClose}>
        <View style={styles.modalContainer}>
          <View style={styles.modalContent}>
            <Text style={styles.label}>Student ID:</Text>
            <TextInput
              placeholder="Type Student ID to check in/out"
              style={styles.input}
            />
            <Button color={Colors.PRIMARY} title="OK" onPress={onClose} />
          </View>
        </View>
      </TouchableWithoutFeedback>
    </Modal>
  );
};

export const Scanner = () => {
  const [hasPermission, setHasPermission] = useState<boolean | null>(null);
  const [scanned, setScanned] = useState(false);
  const [result, setResult] = useState<string | null>(null);
  const [showPopup, setShowPopup] = useState(false);
  const [showInputPopup, setShowInputPopup] = useState(false);
  const [cameraBackground] = useState("transparent");

  const handleBarCodeScanned = ({ data }: { data: string }) => {
    setScanned(true);
    setResult(data);
    setShowPopup(true);
  };

  const handleClosePopup = () => {
    setScanned(false);
    setResult(null);
    setShowPopup(false);
  };

  const handleCloseInputPopup = () => {
    setShowInputPopup(false);
  };

  useEffect(() => {
    (async () => {
      const { status } = await BarCodeScanner.requestPermissionsAsync();
      setHasPermission(status === "granted");
    })();
  }, []);

  if (hasPermission === null) {
    return <Text>Requesting for camera permission</Text>;
  }
  if (hasPermission === false) {
    return <Text>No access to camera</Text>;
  }

  return (
    <View style={styles.container}>
      <View>
        <Header title="Scan" />
      </View>
      <BarCodeScanner
        onBarCodeScanned={scanned ? undefined : handleBarCodeScanned}
        style={styles.camera}
      >
        <View style={styles.cameraMask}>
          <BarcodeMask
            width={300}
            height={300}
            edgeWidth={20}
            edgeHeight={20}
            edgeColor={Colors.PRIMARY}
            edgeRadius={10}
            edgeBorderWidth={5}
            animatedLineColor={Colors.PRIMARY}
            animatedLineHeight={1.5}
            animatedLineWidth={200}
            lineAnimationDuration={2000}
            useNativeDriver
            backgroundColor={cameraBackground}
          />
        </View>
      </BarCodeScanner>
      <View style={styles.inputcontainer}>
        <TouchableOpacity
          onPress={() => setShowInputPopup(true)}
          style={styles.button}
        >
          <Text style={styles.buttontext}>Add Record</Text>
        </TouchableOpacity>
      </View>
      {showInputPopup && (
        <InputPopup
          isVisible={showInputPopup}
          onClose={handleCloseInputPopup}
        />
      )}
      {showPopup && (
        <ResultPopup
          isVisible={showPopup}
          result={result}
          onClose={handleClosePopup}
        />
      )}
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    height: screenHeight,
    backgroundColor: Colors.WHITE,
  },
  modalContainer: {
    flex: 1,
    justifyContent: "center",
    alignItems: "center",
    backgroundColor: "rgba(0,0,0,0.5)",
  },
  camera: {
    height: "70%",
    alignSelf: "center",
    width: screenWidth,
  },
  cameraMask: {
    flex: 1,
  },
  modalContent: {
    backgroundColor: "white",
    padding: 16,
    borderRadius: 8,
    elevation: 5,
    width: 0.8 * screenWidth,
  },
  button: {
    backgroundColor: Colors.PRIMARY,
    padding: 10,
    borderRadius: 5,
  },
  buttontext: {
    color: Colors.WHITE,
  },
  input: {
    padding: 4,
    marginBottom: 12,
    borderWidth: 1,
    borderColor: "gray",
    borderRadius: 5,
  },
  inputcontainer: {
    alignItems: "center",
    justifyContent: "center",
    padding: 5,
    backgroundColor: "red",
    fontFamily: "Poppins_400Regular",
  },
  label: {
    fontFamily: "Poppins_600SemiBold",
  },
});
