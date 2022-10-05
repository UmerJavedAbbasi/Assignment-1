package assignment01bca

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
)


var last_block_hash string
var genesis_block=1

type block struct{
	transaction string
	nonce int
	previousHash string
	currentHash string
}

func (block_list *Block_list)NewBlock(transaction string,nonce int,previousHash string)*block{

	new_block:=new(block)
	new_block.transaction=transaction
	new_block.nonce=nonce
	if genesis_block==1{
		new_block.previousHash=""
		genesis_block=0
	}else{
		new_block.previousHash=last_block_hash
	}
	
//	new_block.currentHash="abc"
	new_block.currentHash=new_block.CreateHash()
	last_block_hash=new_block.currentHash
	block_list.list=append(block_list.list,new_block)	
	return new_block
		
}
type Block_list struct{
	list []*block

}


func (block_list *Block_list)ListBlocks(){

	fmt.Println("List of Blocks")
	for i:=0;i<len(block_list.list);i++{
		fmt.Printf("%s List %d %s \n",strings.Repeat("=",25),i,strings.Repeat("=",25))
		fmt.Printf("Transaction\t%s\n",block_list.list[i].transaction)
		fmt.Printf("Nonce\t%d\n",block_list.list[i].nonce)
		fmt.Printf("Previous Block Hash\t%s\n",block_list.list[i].previousHash)
		fmt.Printf("Current Block Hash\t%s\n",block_list.list[i].currentHash)
	}

}
func (myblock *block)CreateHash()(string){//[32]uint8
	
	curr_block:=(myblock.transaction)+strconv.Itoa(myblock.nonce)+(myblock.previousHash)
	hash := sha256.Sum256([]byte(curr_block))
//	fmt.Printf("%s\t\t%x\t%T\n",curr_block,hash,hash)

	str_hash:=fmt.Sprintf("%x", hash)//converting [32]unit8 to hex string
	//fmt.Printf("String Hash: %s\n",str_hash)
	return str_hash
}
func (block_list *Block_list)ChangeBlock(nonce int,transaction string){

	for i:=0;i<len(block_list.list);i++{
		if block_list.list[i].nonce == nonce{
			block_list.list[i].transaction=transaction
			block_list.list[i].currentHash=block_list.list[i].CreateHash()
			break
		}
	}
}
func (block_list *Block_list)VerifyChain(){
	var blockchain_hash string
	for i:=0;i<len(block_list.list);i++{
		if i!=0{
			block_list.list[i].previousHash=blockchain_hash
		}
		blockchain_hash=block_list.list[i].CreateHash()
	}
	if blockchain_hash==last_block_hash{
		fmt.Println("Blockchain is Not Tempered")
	}else{
		fmt.Println("Blockchain is Tempered")
	}
}
